package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

var body = strings.NewReader(`
{
  "context": {
    "client": {
      "clientName": "ANDROID_EMBEDDED_PLAYER",
      "clientVersion": "16.49"
    }
  },
  "videoId": "HtVdAasjOgU"
}
`)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   val := make(url.Values)
   val["key"] = []string{"AIzaSyCjc_pVEDi4qsv5MtC2dMXzpIaDoRFLsxw"}
   //val["prettyPrint"] = []string{"false"}
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
