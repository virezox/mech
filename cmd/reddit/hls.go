package main

import (
   "fmt"
   "github.com/89z/mech/reddit"
   "github.com/89z/parse/m3u"
   "net/http"
   "os"
   "strconv"
)

func (c choice) HLS(link *reddit.Link) error {
   forms, err := link.HLS()
   if err != nil {
      return err
   }
   for _, form := range forms {
      if c.format {
         fmt.Printf("%+v\n", form)
      } else if c.ids[strconv.Itoa(form.ID)] {
         fmt.Println("GET", form.URI)
         res, err := http.Get(form.URI.String())
         if err != nil {
            return err
         }
         defer res.Body.Close()
         forms, err := m3u.Decode(res.Body, form.URI.Dir)
         if err != nil {
            return err
         }
         for _, form := range forms {
            fmt.Println("GET", form.URI)
            res, err := http.Get(form.URI.String())
            if err != nil {
               return err
            }
            defer res.Body.Close()
            file, err := os.Create(form.URI.File)
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
