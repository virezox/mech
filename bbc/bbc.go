package bbc

import (
   "encoding/json"
   "github.com/89z/format"
   "github.com/89z/format/m3u"
   "net/http"
   "os"
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

var (
   Decode = m3u.Decode
   Log = format.Log{Writer: os.Stdout}
)

type Media struct {
   Kind string
   Type string
   Connection []Video
}

func (m Media) Video() (*Video, error) {
   for _, vid := range m.Connection {
      if vid.Protocol == "http" &&
      vid.Supplier == "mf_akamai" &&
      vid.TransferFormat == "hls" {
         return &vid, nil
      }
   }
   return nil, notFound{"http,hls,mf_akamai"}
}

type NewsVideo struct {
   ExternalID string
   Caption string
}

func NewNewsVideo(addr string) (*NewsVideo, error) {
   req, err := http.NewRequest(
      "GET", producer + path.Base(addr), nil,
   )
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var item newsItem
   if err := json.NewDecoder(res.Body).Decode(&item); err != nil {
      return nil, err
   }
   for _, rel := range item.Relations {
      if rel.PrimaryType == videoType {
         video := new(NewsVideo)
         err := json.Unmarshal(rel.Content, video)
         if err != nil {
            return nil, err
         }
         return video, nil
      }
   }
   return nil, notFound{videoType}
}

func (n NewsVideo) Media() (*Media, error) {
   req, err := http.NewRequest("GET", mediaSelector + n.ExternalID, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var selector struct {
      Media []Media
   }
   if err := json.NewDecoder(res.Body).Decode(&selector); err != nil {
      return nil, err
   }
   for _, media := range selector.Media {
      if media.Kind == "video" {
         return &media, nil
      }
   }
   return nil, notFound{"video"}
}

type Video struct {
   Protocol string
   Supplier string
   TransferFormat string
   Href string
}

func (v Video) HLS() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", v.Href, nil)
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   dir, _ := path.Split(v.Href)
   return Decode(res.Body, dir)
}

type newsItem struct {
   Relations []struct {
      PrimaryType string
      Content json.RawMessage
   }
}

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}

