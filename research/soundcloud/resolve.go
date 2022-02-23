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
   req.URL.Host = "api-mobile.soundcloud.com"
   req.URL.Path = "/resolve"
   val := make(url.Values)
   val["identifier"] = []string{"https://soundcloud.com/pdis_inpartmaint/harold-budd-perhaps-moss"}
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"OAuth 2-276024-1051873921-AceiNB66oXPOz"}
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
