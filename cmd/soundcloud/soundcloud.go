package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/soundcloud"
   "net/http"
   "os"
   "strings"
)

func download(track *soundcloud.Track) error {
   pro, err := track.Progressive()
   if err != nil {
      return err
   }
   fmt.Println("GET", pro.URL)
   res, err := http.Get(pro.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := track.User.Username + "-" + track.Title + ".mp3"
   file, err := os.Create(strings.Map(format.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
