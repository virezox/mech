package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "os"
   "strings"
)

func post(id, name, version string) (*http.Response, error) {
   body := fmt.Sprintf(`
   {
      "videoId": %q, "context": {
         "client": {"clientName": %q, "clientVersion": %q}
      }
   }
   `, id, name, version)
   req, err := http.NewRequest(
      "POST", youtube.PlayerAPI, strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
   req.URL.RawQuery = val.Encode()
   fmt.Println("POST", req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %v", res.Status)
   }
   return res, nil
}

func main() {
   res, err := post("GuAwIatmN3U", "ANDROID", youtube.VersionAndroid)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   f, err := os.Create("file.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   f.ReadFrom(res.Body)
}
