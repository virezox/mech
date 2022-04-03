package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
   "time"
)

type argument struct {
   ext string
   info bool
   sleep time.Duration
}

func (a argument) band(item *bandcamp.Item) error {
   band, err := item.Band()
   if err != nil {
      return err
   }
   for _, item := range band.Discography {
      err := a.tralbum(&item)
      if err != nil {
         return err
      }
   }
   return nil
}

func (a argument) tralbum(item *bandcamp.Item) error {
   tralb, err := item.Tralbum()
   if err != nil {
      return err
   }
   for _, track := range tralb.Tracks {
      if a.info {
         fmt.Println(track)
      } else if track.Streaming_URL != nil {
         fmt.Println("GET", track.Streaming_URL.MP3_128)
         res, err := http.Get(track.Streaming_URL.MP3_128)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         file, err := os.Create(track.Base() + a.ext)
         if err != nil {
            return err
         }
         defer file.Close()
         pro := format.NewProgress(res)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
         time.Sleep(a.sleep)
      }
   }
   return nil
}
