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

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"NBC-Security key=android_nbcuniversal,version=2.4,time=1655588404739,hash=71aea314e10691585473ba09b540ca44eac604f1b445ce115c87a8fe3f22a2cb"}
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
   for _, id := range accountIDs {
      for _, video := range videos {
         v.Mpx.AccountID = id
         req.URL.Path = "/access/vod/nbcuniversal/" + strconv.Itoa(video)
         json.NewEncoder(buf).Encode(v)
         req.Body = io.NopCloser(buf)
         res, err := new(http.Transport).RoundTrip(&req)
         if err != nil {
            panic(err)
         }
         defer res.Body.Close()
         fmt.Println(id, video, res.Status)
         time.Sleep(time.Second)
      }
   }
}
