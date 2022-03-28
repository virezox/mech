package pbs

import (
   "github.com/89z/format"
   "github.com/89z/format/json"
   "net/http"
   "net/url"
)

var LogLevel format.LogLevel

type VideoBridge struct {
   Encodings []string
}

func NewVideoBridge(addr *url.URL) (*VideoBridge, error) {
   req, err := http.NewRequest("GET", addr.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Cookie", "pbsol.station=KERA")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var (
      sep = []byte("\twindow.videoBridge = ")
      video = new(VideoBridge)
   )
   if err := json.Decode(res.Body, sep, video); err != nil {
      return nil, err
   }
   return video, nil
}
