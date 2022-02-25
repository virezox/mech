package bbc

import (
   "encoding/json"
   "net/http"
   "strconv"
   "strings"
)

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

func (n NewsItem) address() (string, error) {
   var buf strings.Builder
   buf.WriteString("http://open.live.bbc.co.uk")
   buf.WriteString("/mediaselector/6/select/version/2.0/mediaset/pc/vpid/")
   for _, rel := range n.Relations {
      if rel.PrimaryType == "bbc.mobile.news.video" {
         buf.WriteString(rel.Content.ExternalID)
         return buf.String(), nil
      }
   }
   return "", notFound{"bbc.mobile.news.video"}
}

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " not found"
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
