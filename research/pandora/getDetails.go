package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   body := strings.NewReader(`{"pandoraId":"TR:1168903"}`)
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "android-tuner.pandora.com"
   req.URL.Path = "/services/json/"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["method"] = []string{"catalog.v4.getDetails"}
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
