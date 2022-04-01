package abc

import (
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strings"
)

var LogLevel format.LogLevel

func Route(addr string) (*http.Response, error) {
   var buf strings.Builder
   buf.WriteString("http://prod.gatekeeper.us-abc.symphony.edgedatg.com")
   buf.WriteString("/api/ws/pluto/v1/layout/route")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Appversion", "10.23.1")
   req.URL.RawQuery = url.Values{
      "brand": {"001"},
      "device": {"031_04"},
      "url": {addr},
   }.Encode()
   LogLevel.Dump(req)
   return new(http.Transport).RoundTrip(req)
}
