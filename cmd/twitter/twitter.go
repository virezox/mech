package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
)

func main() {
   // b
   var statusID int64
   flag.Int64Var(&statusID, "b", 0, "status ID")
   // f
   var bitrate int64
   flag.Int64Var(&bitrate, "f", 2_176_000, "bitrate")
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
   if statusID >= 1 {
      err := statusPath(statusID, bitrate, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func statusPath(statusID, bitrate int64, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   stat, err := guest.Status(statusID)
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
               name, err := variant.Name(stat, statusID)
               if err != nil {
                  return err
               }
               dst, err := os.Create(name)
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


