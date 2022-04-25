package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

/*
set modify_headers '/~q & ~u vod.stream/X-Forwarded-For/25.0.0.0'
set modify_headers '/~u vod.stream/X-Forwarded-For/25.0.0.0'
stream reset by client (PROTOCOL_ERROR)
https://www.channel4.com/programmes/frasier/on-demand/18926-001
*/
func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "www.channel4.com"
   req.URL.Path = "/vod/stream/18926-001"
   req.URL.Scheme = "https"
   req.Header["X-Forwarded-For"] = []string{"25.0.0.0"}
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
