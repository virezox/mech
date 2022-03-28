package pbs

import (
   "github.com/89z/format/json"
   "net/http"
)

type Frontline struct {
   Graph []struct {
      EmbedURL string
   } `json:"@graph"`
}

func NewFrontline(addr string) (*Frontline, error) {
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
      line = new(Frontline)
      sep = []byte(`"application/ld+json">`)
   )
   if err := json.Decode(res.Body, sep, line); err != nil {
      return nil, err
   }
   return line, nil
}
