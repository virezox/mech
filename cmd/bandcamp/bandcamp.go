package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bandcamp"
   "io"
   "net/http"
   "os"
   "time"
)

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
         file, err := os.Create(track.Base() + a.ext)
         if err != nil {
            return err
         }
         pro := format.NewProgress(file, res.ContentLength)
         if _, err := io.Copy(pro, res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
         if err := file.Close(); err != nil {
            return err
         }
         time.Sleep(a.sleep)
      }
   }
   return nil
}

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

