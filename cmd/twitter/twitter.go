package main

import (
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
   "path"
   "strconv"
)

func statusPath(statusID int64, info bool, format int) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   stat, err := twitter.NewStatus(guest, statusID)
   if err != nil {
      return err
   }
   for index, variant := range stat.Variants() {
      addr := variant.URL.String()
      switch {
      case info:
         fmt.Print("ID:", index)
         fmt.Print(" URL:", addr)
         fmt.Println()
      case format == index:
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         name := filename(stat.User.Name, addr, statusID)
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
