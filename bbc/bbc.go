package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "path"
   "strconv"
)

const (
   producer = "http://walter-producer-cdn.api.bbci.co.uk/content/cps/news/"
   videoType = "bbc.mobile.news.video"
)

const mediaSelector =
   "http://open.live.bbc.co.uk" +
   "/mediaselector/6/select/version/2.0/mediaset/pc/vpid/"

var LogLevel format.LogLevel

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

func (m Media) Name(item *NewsItem) (string, error) {
   ext, err := format.ExtensionByType(m.Type)
   if err != nil {
      return "", err
   }
   return item.ShortName + "-" + item.IstatsLabels.CPS_Asset_ID + ext, nil
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

func NewNewsItem(addr string) (*NewsItem, error) {
   req, err := http.NewRequest("GET", producer + path.Base(addr), nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   item := new(NewsItem)
   if err := json.NewDecoder(res.Body).Decode(item); err != nil {
      return nil, err
   }
   return item, nil
}

func (n NewsItem) Media() (*Media, error) {
   var id string
   for _, rel := range n.Relations {
      if rel.PrimaryType == videoType {
         id = rel.Content.ExternalID
      }
   }
   if id == "" {
      return nil, notFound{videoType}
   }
   req, err := http.NewRequest("GET", mediaSelector + id, nil)
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

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
