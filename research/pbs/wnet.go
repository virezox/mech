package pbs

import (
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

type Video struct {
   Full_Length map[string]struct {
      Video_Iframe string
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
      sep = []byte(" STAGE_VIDEOS =")
      vid = new(Video)
   )
   if err := json.Decode(res.Body, sep, vid); err != nil {
      return nil, err
   }
   return vid, nil
}

func (v Video) Widget() (*Widget, error) {
   for _, val := range v.Full_Length {
      addr, err := url.Parse(val.Video_Iframe)
      if err != nil {
         return nil, err
      }
      addr.Scheme = "https"
      return NewWidget(addr)
   }
   return nil, notFound{"full_length"}
}
