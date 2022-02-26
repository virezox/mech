package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "path"
   "strings"
)

var LogLevel format.LogLevel

func NewNewsItem(addr string) (*NewsItem, error) {
   var buf strings.Builder
   buf.WriteString("http://walter-producer-cdn.api.bbci.co.uk")
   buf.WriteString("/content/cps/news/")
   buf.WriteString(path.Base(addr))
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
   item := new(NewsItem)
   if err := json.NewDecoder(res.Body).Decode(item); err != nil {
      return nil, err
   }
   return item, nil
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

func (m Media) href() string {
   for _, video := range m.Connection {
      if video.Protocol == "http" {
         if video.TransferFormat == "hls" {
            if video.Supplier == "mf_akamai" {
               return video.Href
            }
         }
      }
   }
   return ""
}

func (m Media) Name(item *NewsItem) (string, error) {
   ext, err := format.ExtensionByType(m.Type)
   if err != nil {
      return "", err
   }
   return item.ShortName + "-" + item.IstatsLabels.CPS_Asset_ID + ext, nil
}
