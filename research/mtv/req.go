package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "media-utils.mtvnservices.com"
   req.URL.Path = "/services/MediaGenerator/mgid:arc:video:mtv.com:d2eca5e3-a0c9-4058-9fb9-fd7459787e52"
   req.URL.Scheme = "http"
   val := make(url.Values)
   val["acceptMethods"] = []string{"hls"}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
