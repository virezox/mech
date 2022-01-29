package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "path"
   "strconv"
   "strings"
)

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

func (m Media) address() (string, error) {
   for _, video := range m.Connection {
      if video.Protocol == "http" {
         if video.TransferFormat == "hls" {
            if video.Supplier == "mf_akamai" {
               return video.Href, nil
            }
         }
      }
   }
   return "", notFound{"http,hls,mf_akamai"}
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
   var str strings.Builder
   str.WriteString("http://walter-producer-cdn.api.bbci.co.uk")
   str.WriteString("/content/cps/news/")
   str.WriteString(path.Base(addr))
   req, err := http.NewRequest("GET", str.String(), nil)
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

func (n NewsItem) address() (string, error) {
   var addr strings.Builder
   addr.WriteString("http://open.live.bbc.co.uk")
   addr.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   for _, rel := range n.Relations {
      if rel.PrimaryType == "bbc.mobile.news.video" {
         addr.WriteString(rel.Content.ExternalID)
         return addr.String(), nil
      }
   }
   return "", notFound{"bbc.mobile.news.video"}
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
