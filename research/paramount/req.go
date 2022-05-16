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
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "cbsi.live.ott.irdeto.com"
   req.URL.Path = "/widevine/getlicense"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["AccountId"] = []string{"cbsi"}
   val["ContentId"] = []string{"eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"}
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

var body = strings.NewReader(``)
