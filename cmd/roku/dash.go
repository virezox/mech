package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech"
   "github.com/89z/mech/roku"
   "io"
   "net/http"
   "os"
)

func doDASH(con *roku.Content, bandwidth int, info bool) error {
   video := con.DASH()
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   adas, err := dash.NewAdaptationSet(res.Body)
   if err != nil {
      return err
   }
   if info {
      for _, ada := range adas.MimeType(dash.Video) {
         for _, rep := range ada.Representation {
            fmt.Println(rep)
         }
      }
   } else {
      video := adas.MimeType(dash.Video).Represent(0)
      addrs, err := video.URL(res.Request.URL)
      if err != nil {
         return err
      }
      ext, err := mech.ExtensionByType(dash.Video)
      if err != nil {
         return err
      }
      file, err := os.Create(con.Base() + ext)
      if err != nil {
         return err
      }
      defer file.Close()
      pro := format.ProgressChunks(file, len(addrs))
      for _, addr := range addrs {
         res, err := http.Get(addr.String())
         if err != nil {
            return err
         }
         pro.AddChunk(res.ContentLength)
         if _, err := io.Copy(pro, res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
