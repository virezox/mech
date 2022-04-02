package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech"
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
      } else if track.Streaming_URL != nil {
         fmt.Println("GET", track.Streaming_URL.MP3_128)
         res, err := http.Get(track.Streaming_URL.MP3_128)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         ext, err := mech.Ext(res.Header)
         if err != nil {
            return err
         }
         file, err := os.Create(tralb.Base() + ext)
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
   return nil
}
