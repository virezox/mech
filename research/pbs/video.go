package pbs

import (
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

type Content struct {
   ContentURL string
   Video struct {
      ContentURL string
   }
}

func NewContent(addr string) (*Content, error) {
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
      con = new(Content)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Decode(res.Body, sep, con); err != nil {
      return nil, err
   }
   return con, nil
}

func (c Content) Widget() (*Widget, error) {
   if c.ContentURL == "" {
      c.ContentURL = c.Video.ContentURL
   }
   addr, err := url.Parse(c.ContentURL)
   if err != nil {
      return nil, err
   }
   return NewWidget(addr)
}
