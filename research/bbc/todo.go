package bbc

import (
   "encoding/json"
   "net/http"
   "strings"
)

type Media struct {
   Kind string
   Type string
   Connection []Connection
}

func (m Media) GetConnection() *Connection {
   for _, con := range m.Connection {
      if con.Protocol == "http" {
         if con.Supplier == "mf_akamai" {
            if con.TransferFormat == "hls" {
               return &con
            }
         }
      }
   }
   return nil
}

type Mediaset struct {
   Media []Media
}

func (m Mediaset) GetMedia() *Media {
   for _, med := range m.Media {
      if med.Kind == "video" {
         return &med
      }
   }
   return nil
}

func (n NewsItem) Relation() *Relation {
   for _, rel := range n.Relations {
      if rel.PrimaryType == "bbc.mobile.news.video" {
         return &rel
      }
   }
   return nil
}

type Relation struct {
   PrimaryType string
   Content struct {
      ExternalID string
   }
}

func (r Relation) Mediaset() (*Mediaset, error) {
   var buf strings.Builder
   buf.WriteString("http://open.live.bbc.co.uk")
   buf.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   buf.WriteString(r.Content.ExternalID)
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   set := new(Mediaset)
   if err := json.NewDecoder(res.Body).Decode(set); err != nil {
      return nil, err
   }
   return set, nil
}
