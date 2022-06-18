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
   "device": "android",
   "deviceId": "52966deadb9df95",
   "externalAdvertiserId": "TELE_VOD_9000245869",
   "mpx": {
      "accountId": "2304991196"
   }
}
`)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"NBC-Security key=android_nbcuniversal,version=2.4,hash=278d8ac462bb0c4e8bea8240871b9a880b4ec4edd40d77bfb14485df72b12974,time=1655574676072"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "access-cloudpath.media.nbcuni.com"
   req.URL.Path = "/access/vod/nbcuniversal/9000221348"
   req.URL.Scheme = "http"
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
