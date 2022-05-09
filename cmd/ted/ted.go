package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/ted"
   "io"
   "net/http"
   "os"
)

func process(slug string, info bool, bitrate int64) error {
   talk, err := ted.NewTalkResponse(slug)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(talk)
   } else {
      video := talk.Video(bitrate)
      fmt.Println("GET", video.URL)
      res, err := http.Get(video.URL)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      file, err := os.Create(slug + "." + video.Format)
      if err != nil {
         return err
      }
      defer file.Close()
      pro := format.ProgressBytes(file, res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
   }
   return nil
}
