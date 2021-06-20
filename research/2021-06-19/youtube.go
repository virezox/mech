package main

import (
   "fmt"
   "net/http"
   "os"
   "strings"
)

func request() error {
   payload := `
   {
      "videoId":"dQw4w9WgXcQ",
      "context": {
         "client": {"hl":"en","gl":"US","clientName":"ANDROID","clientVersion":"16.02"}
      }
   }
   `
   req, err := http.NewRequest(
      "POST", "https://youtubei.googleapis.com/youtubei/v1/player",
      strings.NewReader(payload),
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Content-Type", "application/json")
   req.Header.Set("User-Agent", "com.google.android.youtube/16.02.35(Linux; U; Android 10; en_US; Pixel 4 XL Build/QQ3A.200805.001) gzip")
   req.Header.Set("x-goog-api-format-version", "2")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return fmt.Errorf("status %v", res.Status)
   }
   f, err := os.Create("file.json")
   if err != nil {
      return err
   }
   defer f.Close()
   f.ReadFrom(res.Body)
   return nil
}

func main() {
   err := request()
   if err != nil {
      panic(err)
   }
}
