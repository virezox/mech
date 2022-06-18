package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

type vodRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID int `json:"accountId"`
   } `json:"mpx"`
}

var videos = []int{
   // nbc.com/botched/video/seeing-double/3049418
   // 2304982139 3049418 200 OK
   // "resourceID": "e",
   3049418,
   // nbc.com/pasion-de-gavilanes/video/la-valentia-de-norma/9000221348
   // 2304991196 9000221348 200 OK
   // "resourceID": "telemundo",
   9000221348,
}

var accounts = []int{
   2304982139,
   2304991196,
}

const auth = "NBC-Security key=android_nbcuniversal,version=2.4,time=1655590454716,hash=ad3659a765dc2e8ceb670f9dadea75e0f2f9f012ffaa399051fc20394297b0ce"

func main() {
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
   for _, account := range accounts {
      for _, video := range videos {
         req.URL.Path = "/access/vod/nbcuniversal/" + strconv.Itoa(video)
         v.Mpx.AccountID = account
         json.NewEncoder(buf).Encode(v)
         req.Body = io.NopCloser(buf)
         res, err := new(http.Transport).RoundTrip(&req)
         if err != nil {
            panic(err)
         }
         defer res.Body.Close()
         fmt.Println(v.Mpx.AccountID, video, res.Status)
         time.Sleep(time.Second)
      }
   }
}
