package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "strings"
   "time"
)

type client struct {
   name string
   version string
}

var clients = []client{
   {"ANDROID", "16.07.34"},
   {"ANDROID_CREATOR", "21.06.103"},
   {"ANDROID_KIDS", "6.02.3"},
   {"ANDROID_MUSIC", "4.16.51"},
   {"IOS", "16.05.7"},
   {"IOS_CREATOR", "20.47.100"},
   {"IOS_KIDS", "5.42.2"},
   {"IOS_MUSIC", "4.16.1"},
   {"TVHTML5", "7.20210224.00.00"},
   {"WEB", "2.20210223.09.00"},
   {"WEB_CREATOR", "1.20210223.01.00"},
   {"WEB_KIDS", "2.1.3"},
   {"WEB_REMIX", "0.1"},
}

func request() error {
   for _, c := range clients {
      payload := fmt.Sprintf(`
      {
         "videoId":"dQw4w9WgXcQ",
         "context": {
            "client": {
               "clientName": %q,
               "clientVersion": %q
            }
         }
      }
      `, c.name, c.version)
      req, err := http.NewRequest(
         "POST",
         "https://www.youtube.com/youtubei/v1/player",
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
      body, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      fmt.Print(c)
      if bytes.Contains(body, []byte("publishDate")) {
         fmt.Println("pass")
      } else {
         fmt.Println("fail")
      }
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}

func main() {
   err := request()
   if err != nil {
      panic(err)
   }
}
