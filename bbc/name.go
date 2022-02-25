package bbc

import (
   "github.com/89z/format"
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

func (n NewsItem) Media() (*Media, error) {
   addr, err := n.address()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("GET", addr, nil)
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
   for _, media := range mediaset.Media {
      if media.Kind == "video" {
         return &media, nil
      }
   }
   return nil, notFound{"video"}
}

func (m Media) Name(item *NewsItem) (string, error) {
   ext, err := format.ExtensionByType(m.Type)
   if err != nil {
      return "", err
   }
   return item.ShortName + "-" + item.IstatsLabels.CPS_Asset_ID + ext, nil
}
