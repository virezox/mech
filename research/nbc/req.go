package main

import (
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Body = io.NopCloser(req_body)
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      buf, err := httputil.DumpResponse(res, true)
      if err != nil {
         panic(err)
      }
      os.Stdout.Write(buf)
   }
   raw_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   body := string(raw_body)
   fmt.Println(body)
   if strings.Contains(body, `"convivaAssetName"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
   if strings.Contains(body, `"mpxAccountId"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
   if strings.Contains(body, `"NBCE680517903"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

var req_body = strings.NewReader(fmt.Sprintf(`
{
  "operationName": "bonanzaPage",
  "query": %q,
  "variables": {
    "app": "nbc",
    "appVersion": 7033000,
    "authorized": false,
    "brand": null,
    "device": "android",
    "endCardMpxGuid": null,
    "endCardTagLine": null,
    "isDayZero": true,
    "language": "en",
    "ld": true,
    "minimumTiles": -1,
    "mpxGuid": null,
    "name": "NBCE680517903",
    "nationalBroadcastType": "",
    "nbcAffiliateName": "",
    "categories": null,
    "oneApp": true,
    "platform": "android",
    "playlistMachineName": null,
    "profile": [
      "00000",
      "11111"
    ],
    "queryName": null,
    "telemundoAffiliateName": "",
    "timeZone": "America/Chicago",
    "type": "VIDEO",
    "userId": "-3379434425322817474"
  }
}
`, query))
