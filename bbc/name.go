package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
)

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

func (m Media) Name(item *NewsItem) (string, error) {
   ext, err := format.ExtensionByType(m.Type)
   if err != nil {
      return "", err
   }
   return item.ShortName + "-" + item.IstatsLabels.CPS_Asset_ID + ext, nil
}

func (n NewsItem) Media() (*Media, error) {
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
   var mediaset struct {
      Media []Media
   }
   if err := json.NewDecoder(res.Body).Decode(&mediaset); err != nil {
      return nil, err
   }
   var media Media
   for _, media = range mediaset.Media {
      if media.Kind == "video" {
         break
      }
   }
   return &media, nil
}

type Media struct {
   Kind string
   Type string
   Connection []struct {
      Protocol string
      Supplier string
      TransferFormat string
      Href string
   }
}
