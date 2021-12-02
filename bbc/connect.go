package bbc

import (
   "encoding/json"
   "encoding/xml"
   "github.com/89z/mech"
   "net/http"
)

func (n NewsVideo) Connection() (*Connection, error) {
   req, err := http.NewRequest("GET", mediaSelector + n.ExternalID, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var sel selector
   if err := json.NewDecoder(res.Body).Decode(&sel); err != nil {
      return nil, err
   }
   for _, media := range sel.Media {
      if media.Kind == "video" {
         var cons []Connection
         err := json.Unmarshal(media.Connection, &cons)
         if err != nil {
            return nil, err
         }
         for _, con := range cons {
            if con.Protocol == "http" && con.Supplier == "mf_akamai" {
               if con.TransferFormat == "dash" {
                  return &con, nil
               }
            }
         }
         return nil, mech.NotFound{"http,dash,mf_akamai"}
      }
   }
   return nil, mech.NotFound{"video"}
}

func (c Connection) MPD() (*MPD, error) {
   req, err := http.NewRequest("GET", c.Href, nil)
   if err != nil {
      return nil, err
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   mpd := new(MPD)
   if err := xml.NewDecoder(res.Body).Decode(mpd); err != nil {
      return nil, err
   }
   return mpd, nil
}

type MPD struct {
   Period struct {
      AdaptationSet []struct {
         MimeType string `xml:"mimeType,attr"`
         Representation struct {
            Height int `xml:"height,attr"`
         }
      }
   }
}
