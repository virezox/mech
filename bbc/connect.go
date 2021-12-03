package bbc

import (
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
   "sort"
)

func (c Connection) HLS() ([]HLS, error) {
   req, err := http.NewRequest("GET", c.Href, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prefix, _ := path.Split(c.Href)
   var hlss []HLS
   for key, val := range m3u.NewPlaylist(res.Body) {
      hlss = append(hlss, HLS{
         Resolution: val["RESOLUTION"], URI: prefix + key,
      })
   }
   sort.Slice(hlss, func(a, b int) bool {
      return hlss[a].Resolution < hlss[b].Resolution
   })
   for key := range hlss {
      hlss[key].ID = key
   }
   return hlss, nil
}

type Connection struct {
   Protocol string
   Supplier string
   TransferFormat string
   Href string
}

type HLS struct {
   ID int
   Resolution string
   URI string
}


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
               if con.TransferFormat == "hls" {
                  return &con, nil
               }
            }
         }
         return nil, mech.NotFound{"http,hls,mf_akamai"}
      }
   }
   return nil, mech.NotFound{"video"}
}
