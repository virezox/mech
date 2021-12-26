package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/reddit"
   "github.com/89z/parse/m3u"
   "net/http"
   "os"
   "path"
   "strconv"
   "strings"
)

type choice struct {
   info bool
   formats map[string]bool
}

func (c choice) DASH(link *reddit.Link) error {
   dash, err := link.DASH()
   if err != nil {
      return err
   }
   for _, ada := range dash.Period.AdaptationSet {
      for _, rep := range ada.Representation {
         if c.info {
            fmt.Printf("%+v\n", rep)
         } else if c.formats[rep.ID] {
            res, err := http.Get(rep.BaseURL)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            ext, err := mech.ExtensionByType(rep.MimeType)
            if err != nil {
               return err
            }
            name := link.Subreddit + "-" + link.Title + ext
            file, err := os.Create(strings.Map(mech.Clean, name))
            if err != nil {
               return err
            }
            defer file.Close()
            file.ReadFrom(res.Body)
         }
      }
   }
   return nil
}

func (c choice) HLS(link *reddit.Link) error {
   forms, err := link.HLS()
   if err != nil {
      return err
   }
   for id, form := range forms {
      addr := form["URI"]
      switch {
      case c.info && !strings.Contains(addr, "_iframe."):
         fmt.Println(id, form)
      case c.formats[strconv.Itoa(id)]:
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
