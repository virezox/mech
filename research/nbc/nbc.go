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
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Host"] = []string{"friendship.nbc.co"}
   req.Header["User-Agent"] = []string{"okhttp/4.6.0"}
   req.Method = "GET"
   req.URL = new(url.URL)
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["extensions"] = []string{"{\"persistedQuery\":{\"sha256Hash\":\"8100bcb54b4f5b311588ec99ae5f4c4f9acd1fac6ec7709d24642778118491c5\",\"version\":1}}"}
   val["variables"] = []string{`
   {
      "app": "nbc",
      "appVersion": 7031000,
      "authorized": false,
      "device": "android",
      "language": "en",
      "ld": true,
      "minimumTiles": -1,
      "name": "9000221348",
      "nationalBroadcastType": "",
      "nbcAffiliateName": "",
      "oneApp": true,
      "platform": "android",
      "profile": [
      "00000",
      "11111"
      ],
      "telemundoAffiliateName": "",
      "timeZone": "America/Chicago",
      "type": "VIDEO",
      "userId": "8292284999374523746"
   }
   `}
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
