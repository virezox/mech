package main

import (
   "fmt"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
   "path"
   "strings"
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
         name := filename(stat.User.Name, id, addr)
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

func filename(name, id, addr string) string {
   var buf strings.Builder
   buf.WriteString(name)
   buf.WriteByte('-')
   buf.WriteString(id)
   buf.WriteString(path.Ext(addr))
   return buf.String()
}
