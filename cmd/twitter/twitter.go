package main

import (
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
)

func statusPath(statusID, bitrate int64, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   stat, err := twitter.NewStatus(guest, statusID)
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
