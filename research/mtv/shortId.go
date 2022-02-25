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
   req.URL.Host = "neutron-api.viacom.tech"
   req.URL.Path = "/api/2.8/property"
   val := make(url.Values)
   req.URL.Scheme = "https"
   val["brand"] = []string{"mtv"}
   val["platform"] = []string{"web"}
   val["region"] = []string{"US"}
   val["shortId"] = []string{"9i96kj"}
   val["type"] = []string{"episode"}
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
