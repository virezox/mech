package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
   "time"
)

func doBand(item *bandcamp.Item, info bool, sleep time.Duration) error {
   band, err := item.Band()
   if err != nil {
      return err
   }
   for _, item := range band.Discography {
      err := doTralbum(&item, info, sleep)
      if err != nil {
         return err
      }
   }
   return nil
}

func doTralbum(item *bandcamp.Item, info bool, sleep time.Duration) error {
   tralb, err := item.Tralbum()
   if err != nil {
      return err
   }
   for _, track := range tralb.Tracks {
      if info {
         fmt.Printf("%+v\n", track)
      } else {
         addr, ok := track.MP3_128()
         if ok {
            fmt.Println("GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            name, err := track.Name(tralb, res.Header)
            if err != nil {
               return err
            }
            file, err := os.Create(name)
            if err != nil {
               return err
            }
            defer file.Close()
            pro := format.NewProgress(res)
            if _, err := file.ReadFrom(pro); err != nil {
               return err
            }
            time.Sleep(sleep)
         }
      }
   }
   return nil
}
