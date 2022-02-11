package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
   "strings"
   "time"
)

type flags struct {
   info bool
   sleep time.Duration
}

func (f flags) process(data *bandcamp.DataTralbum) error {
   for _, track := range data.TrackInfo {
      if f.info {
         fmt.Printf("%+v\n", track)
      } else {
         addr, ok := track.File.MP3_128()
         if ok {
            fmt.Println("GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            name := data.Artist + "-" + track.Title + ".mp3"
            file, err := os.Create(strings.Map(format.Clean, name))
            if err != nil {
               return err
            }
            defer file.Close()
            if _, err := file.ReadFrom(res.Body); err != nil {
               return err
            }
            time.Sleep(f.sleep)
         }
      }
   }
   return nil
}
