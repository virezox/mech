package nbc

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type vodRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID int `json:"accountId"`
   } `json:"mpx"`
}

func access() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{auth}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "access-cloudpath.media.nbcuni.com"
   req.URL.Scheme = "http"
   var v vodRequest
   v.Device = "android"
   v.DeviceID = "android"
   v.ExternalAdvertiserID = "NBC"
   buf := new(bytes.Buffer)
   req.URL.Path = "/access/vod/nbcuniversal/" + strconv.Itoa(video)
   v.Mpx.AccountID = account
   json.NewEncoder(buf).Encode(v)
   req.Body = io.NopCloser(buf)
   return new(http.Transport).RoundTrip(&req)
}

func friendship() (*http.Response, error) {
   var body = strings.NewReader(`
   {
      "variables": {
         "app": "nbc",
         "name": "9000221348",
         "platform": "android",
         "type": "VIDEO",
         "userId": "",
         "oneApp": true
      },
      "extensions": {
         "persistedQuery": {
            "sha256Hash": "872a3dffc3ae6cdb3dc69fe3d9a949b539de7b579e95b2942e68d827b1a6ec62"
         }
      }
   }
   `)
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header.Set("Content-Type", "application/json")
   req.URL = new(url.URL)
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
   req.Method = "POST"
   req.URL.Scheme = "https"
   return new(http.Transport).RoundTrip(&req)
}
