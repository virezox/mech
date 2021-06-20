package main

import (
   "fmt"
   "os"
   "net/http"
   "strings"
)

func request() error {
   payload := `
   {
      "videoId":"dQw4w9WgXcQ",
      "context": {
         "client": {
            "clientName": "WEB",
            "clientVersion": "2.20210223.09.00"
         }
      }
   }
   `
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      strings.NewReader(payload),
   )
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      fmt.Println(payload)
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
