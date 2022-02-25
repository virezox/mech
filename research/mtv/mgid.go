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
   req.URL.Path = "/api/2.9/property/mgid:arc:episode:mtv.com:97152590-d238-11e1-a549-0026b9414f30"
   val := make(url.Values)
   val["brand"] = []string{"mtv"}
   val["platform"] = []string{"android"}
   val["region"] = []string{"US"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
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
