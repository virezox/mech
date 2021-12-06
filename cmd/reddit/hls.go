package main

import (
   "fmt"
   "github.com/89z/mech/reddit"
   "github.com/89z/parse/m3u"
   "net/http"
   "os"
   "path"
   "strconv"
   "strings"
)

func (c choice) HLS(link *reddit.Link) error {
   forms, err := link.HLS()
   if err != nil {
      return err
   }
   for id, form := range forms {
      addr := form["URI"]
      switch {
      case c.format && !strings.Contains(addr, "_iframe."):
         fmt.Println(id, form)
      case c.ids[strconv.Itoa(id)]:
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         dir, _ := path.Split(addr)
         forms, err := m3u.Decode(res.Body, dir)
         if err != nil {
            return err
         }
         for _, form := range forms {
            addr := form["URI"]
            fmt.Println("GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            _, name := path.Split(addr)
            file, err := os.Create(name)
            if err != nil {
               return err
            }
            defer file.Close()
            if _, err := file.ReadFrom(res.Body); err != nil {
               return err
            }
         }
      }
   }
   return nil
}
