package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/facebook"
   "io"
   "net/http"
   "net/url"
   "os"
   "path"
)

func doVideo(videoID int64, info bool) error {
   video, err := facebook.NewVideo(videoID)
   if err != nil {
      return err
   }
   if info {
      err := video.Preferred_Thumbnail.Parse()
      if err != nil {
         return err
      }
      fmt.Println(video)
   } else {
      addr, err := url.Parse(video.Playable_URL_Quality_HD)
      if err != nil {
         return err
      }
      fmt.Println("GET", addr)
      res, err := http.Get(addr.String())
      if err != nil {
         return err
      }
      defer res.Body.Close()
      file, err := os.Create(path.Base(addr.Path))
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
