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
   req.URL.Host = "www.instagram.com"
   req.URL.Path = "/p/CLHoAQpCI2i/"
   val := make(url.Values)
   val["__a"] = []string{"1"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Header.Set("Cookie", "sessionid=" + sessionID)
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
