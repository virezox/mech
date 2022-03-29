package pbs

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

var LogLevel format.LogLevel

type Widget struct {
   Encodings []string
}

func NewWidget(addr *url.URL) (*Widget, error) {
   req, err := http.NewRequest("GET", addr.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Cookie", "pbsol.station=KERA")
   LogLevel.Dump(req)
   // this can redirect
   res, err := new(http.Client).Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var (
      sep = []byte("\twindow.videoBridge = ")
      wid = new(Widget)
   )
   if err := json.Decode(res.Body, sep, wid); err != nil {
      return nil, err
   }
   return wid, nil
}
