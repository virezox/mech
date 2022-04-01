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
   req.URL.Host = "prod.gatekeeper.us-abc.symphony.edgedatg.com"
   req.URL.Path = "/api/ws/pluto/v1/layout/route"
   val := make(url.Values)
   val["url"] = []string{"/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you"}
   val["brand"] = []string{"001"}
   val["device"] = []string{"031_04"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "http"
   req.Header["Appversion"] = []string{"10.23.1"}
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
