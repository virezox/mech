package main

import (
   "bytes"
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "strings"
)

const PlayerAPI = "https://www.youtube.com/youtubei/v1/player"

/*
web normal
62870
59609
avg 61239.5

web gzip
20488
20474
avg 20481
66% reduce

android normal
59133
57675
avg 58404

android gzip
12614
12717
12665.5
78%
*/

func request() error {
   body := fmt.Sprintf(`
   {
      "videoId": "Ht5d7gVqo1I", "context": {
         "client": {"clientName": "ANDROID", "clientVersion": %q}
      }
   }
   `, youtube.VersionAndroid)
   req, err := http.NewRequest(
      "POST", PlayerAPI, strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Accept-Encoding", "gzip")
   fmt.Println("POST", req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   buf := new(bytes.Buffer)
   buf.ReadFrom(res.Body)
   println(buf.Len())
   return nil
}

func main() {
   err := request()
   if err != nil {
      panic(err)
   }
}
