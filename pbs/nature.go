package pbs

import (
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

func NewNature(addr string) (*Nature, error) {
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
   scan, err := json.NewScanner(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Split = []byte(`{"preview"`)
   scan.Scan()
   nat := new(Nature)
   if err := scan.Decode(nat); err != nil {
      return nil, err
   }
   return nat, nil
}

type Nature struct {
   Full_Length map[string]struct {
      Video_Iframe string
   }
}

func (n Nature) Widget() (*Widget, error) {
   for _, full := range n.Full_Length {
      addr, err := url.Parse(full.Video_Iframe)
      if err != nil {
         return nil, err
      }
      addr.Scheme = "https"
      return NewWidget(addr)
   }
   return nil, notFound{"video_iframe"}
}
