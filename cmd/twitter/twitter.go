package main

import (
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
   "path"
   "strconv"
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
   for _, variant := range stat.Variants() {
      switch {
      case info:
         fmt.Println(variant)
      case variant.Bitrate == bitrate:
         fmt.Println("GET", variant.URL)
         res, err := http.Get(variant.URL)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         name := filename(stat.User.Name, variant.URL, statusID)
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
   return nil
}

func filename(name, addr string, id int64) string {
   buf := []byte(name)
   buf = append(buf, '-')
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, path.Ext(addr)...)
   return string(buf)
}
