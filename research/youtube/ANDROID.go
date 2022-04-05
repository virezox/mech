package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
   gp "github.com/89z/googleplay"
)

func clientVersion() (string, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return "", err
   }
   token, err := gp.OpenToken(cache, "googleplay/token.json")
   if err != nil {
      return "", err
   }
   device, err := gp.OpenDevice(cache, "googleplay/device.json")
   if err != nil {
      return "", err
   }
   head, err := token.Header(device)
   if err != nil {
      return "", err
   }
   det, err := head.Details("com.google.android.youtube")
   if err != nil {
      return "", err
   }
   return string(det.VersionString), nil
}

func main() {
   var (
      err error
      play player
   )
   play.VideoID = "eZHsmb4ezEk"
   play.Context.Client.ClientName = "ANDROID"
   play.Context.Client.ClientVersion, err = clientVersion()
   if err != nil {
      panic(err)
   }
   buf := new(bytes.Buffer)
   json.NewEncoder(buf).Encode(play)
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player", buf,
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("X-Goog-Api-Key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, play.Context.Client)
}

type player struct {
   VideoID string `json:"videoId"`
   Context struct {
      Client struct {
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
}
