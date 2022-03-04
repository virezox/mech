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
   req.URL.Host = "www.paramountplus.com"
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU.json"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["at"] = []string{"ABAJi4xSDPXIEUKTlJ6BFQpMdL3hrvn5xbm+Xly+9QZJFycgSL/4/YiDMKY4XWomRkI="}
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
