package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/twitter"
   "io"
   "net/http"
   "os"
)

func main() {
   // b
   var status int64
   flag.Int64Var(&status, "b", 0, "status ID")
   // c
   var space string
   flag.StringVar(&space, "c", "", "space ID")
   // f
   var bitrate int64
   flag.Int64Var(&bitrate, "f", 2_176_000, "status bitrate")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      twitter.LogLevel = 1
   }
   if status >= 1 {
      err := doStatus(status, bitrate, info)
      if err != nil {
         panic(err)
      }
   } else if space != "" {
      err := doSpace(space, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
func doSpace(id string, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   space, err := guest.AudioSpace(id)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(space)
   } else {
      source, err := guest.Source(space)
      if err != nil {
         return err
      }
      fmt.Println("GET", source.Location)
      res, err := http.Get(source.Location)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      seg, err := hls.NewScanner(res.Body).Segment(res.Request.URL)
      if err != nil {
         return err
      }
      file, err := os.Create(space.Base() + seg.Ext())
      if err != nil {
         return err
      }
      defer file.Close()
      pro := format.ProgressChunks(file, len(seg.Info))
      for _, info := range seg.Info {
         res, err := http.Get(info.URI.String())
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

func doStatus(id, bitrate int64, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   stat, err := guest.Status(id)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(stat)
   } else {
      for _, media := range stat.Extended_Entities.Media {
         for _, variant := range media.Variants() {
            if variant.Bitrate == bitrate {
               fmt.Println("GET", variant.URL)
               res, err := http.Get(variant.URL)
               if err != nil {
                  return err
               }
               defer res.Body.Close()
               ext, err := variant.Ext()
               if err != nil {
                  return err
               }
               dst, err := os.Create(stat.Base(id) + ext)
               if err != nil {
                  return err
               }
               defer dst.Close()
               if _, err := dst.ReadFrom(res.Body); err != nil {
                  return err
               }
            }
         }
      }
   }
   return nil
}
