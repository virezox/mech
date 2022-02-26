package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

var LogLevel format.LogLevel

func (n NewsItem) address() string {
   var buf strings.Builder
   buf.WriteString("http://open.live.bbc.co.uk")
   buf.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   for _, rel := range n.Relations {
      if rel.PrimaryType == "bbc.mobile.news.video" {
         buf.WriteString(rel.Content.ExternalID)
         return buf.String()
      }
   }
   return ""
}

type NewsItem struct {
   ShortName string
   IstatsLabels struct {
      CPS_Asset_ID string
   }
   Relations []struct {
      PrimaryType string
      Content struct {
         ExternalID string
      }
   }
}

func (n NewsItem) Mediaset() (*Mediaset, error) {
   req, err := http.NewRequest("GET", n.address(), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   set := Mediaset{Parent: &n}
   if err := json.NewDecoder(res.Body).Decode(&set.Child); err != nil {
      return nil, err
   }
   return &set, nil
}

type Mediaset struct {
   Parent *NewsItem
   Child struct {
      Media []struct {
         Kind string
         Type string
         Connection []struct {
            Protocol string
            Supplier string
            TransferFormat string
            Href string
         }
      }
   }
}
