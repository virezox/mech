package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech"
   "github.com/89z/mech/pbs"
   "net/http"
   "os"
)

func download(title string, video pbs.AssetVideo) error {
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := title + "-" + video.Profile + ".mp4"
   file, err := os.Create(mech.Clean(name))
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.NewProgress(res)
   if _, err := file.ReadFrom(pro); err != nil {
      return err
   }
   return nil
}
