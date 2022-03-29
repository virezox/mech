package pbs

import (
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

type Video struct {
   ContentURL string
   Video struct {
      ContentURL string
   }
}

func NewVideo(addr string) (*Video, error) {
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
   var (
      vid = new(Video)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Decode(res.Body, sep, vid); err != nil {
      return nil, err
   }
   return vid, nil
}

func (v Video) Widget() (*Widget, error) {
   if v.ContentURL == "" {
      v.ContentURL = v.Video.ContentURL
   }
   addr, err := url.Parse(v.ContentURL)
   if err != nil {
      return nil, err
   }
   return NewWidget(addr)
}
