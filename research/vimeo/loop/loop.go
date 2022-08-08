package main

import (
   "bytes"
   "errors"
   "github.com/89z/rosso/http"
   "io"
   "strconv"
   "time"
)

const ott_data = ".OTTData "

var client = http.Default_Client

func get(sub string, video int64) error {
   b := []byte("http://embed.vhx.tv/")
   b = append(b, sub...)
   b = append(b, '/')
   b = strconv.AppendInt(b, video, 10)
   b = append(b, "?vimeo=1"...)
   res, err := client.Get(string(b))
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode == 429 {
      panic(res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   if !bytes.Contains(body, []byte(ott_data)) {
      return errors.New(ott_data)
   }
   return nil
}

func main() {
   var video int64 = 17863
   for video <= 28599 {
      err := get("subscriptions", video)
      if err == nil {
         err := get("videos", video)
         if err == nil {
            break
         }
      }
      time.Sleep(199 * time.Millisecond)
      video++
   }
}
