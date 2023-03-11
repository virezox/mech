package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech"
   "github.com/89z/mech/bandcamp"
   "io"
   "net/http"
   "os"
   "time"
)

func doBand(param *bandcamp.Params, info bool, sleep time.Duration) error {
   band, err := param.Band()
   if err != nil {
      return err
   }
   for _, item := range band.Discography {
      tralb, err := item.Tralbum()
      if err != nil {
         return err
      }
      if err := doTralbum(tralb, info, sleep); err != nil {
         return err
      }
   }
   return nil
}

func doTralbum(tralb *bandcamp.Tralbum, info bool, sleep time.Duration) error {
   for _, track := range tralb.Tracks {
      if info {
         fmt.Println(track)
      } else if track.Streaming_URL != nil {
         fmt.Println("GET", track.Streaming_URL.MP3_128)
         res, err := http.Get(track.Streaming_URL.MP3_128)
         if err != nil {
            return err
         }
         ext, err := mech.ExtensionByType(res.Header.Get("Content-Type"))
         if err != nil {
            return err
         }
         file, err := os.Create(track.Base() + ext)
         if err != nil {
            return err
         }
         pro := format.ProgressBytes(file, res.ContentLength)
         if _, err := io.Copy(pro, res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
         if err := file.Close(); err != nil {
            return err
         }
         time.Sleep(sleep)
      }
   }
   return nil
}
