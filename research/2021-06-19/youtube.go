package main

import (
   "fmt"
   "net/http"
   "os"
   "strings"
)

func request() error {
   body := fmt.Sprintf(`
   {
      "videoId": %q, "context": {
         "client": {"clientName": "ANDROID", "clientVersion": "15.01"}
      }
   }
   `, "dQw4w9WgXcQ")
   req, err := http.NewRequest(
      "POST", "https://youtubei.googleapis.com/youtubei/v1/player",
      strings.NewReader(body),
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
