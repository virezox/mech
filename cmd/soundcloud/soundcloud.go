package main

import (
   "fmt"
   "github.com/89z/mech/soundcloud"
   "net/http"
   "os"
)

func download(track *soundcloud.Track) error {
   media, err := track.Progressive()
   if err != nil {
      return err
   }
   fmt.Println("GET", media.URL)
   res, err := http.Get(media.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name, err := media.Name(track)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
