package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/reddit"
   "net/http"
   "os"
   "strings"
)

type choice struct {
   format bool
   ids map[string]bool
}

func (c choice) HLS(link *reddit.Link) error {
   hlss, err := link.HLS()
   if err != nil {
      return err
   }
   for _, hls := range hlss {
      if c.format {
         fmt.Printf("%+v\n", hls)
      } else {
      }
   }
   return nil
}

func (c choice) DASH(link *reddit.Link) error {
   dash, err := link.DASH()
   if err != nil {
      return err
   }
   for _, ada := range dash.Period.AdaptationSet {
      for _, rep := range ada.Representation {
         if c.format {
            fmt.Printf("%+v\n", rep)
         } else if c.ids[rep.ID] {
            res, err := http.Get(rep.BaseURL)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            exts, err := mech.ExtensionsByType(rep.MimeType)
            if exts == nil {
               return fmt.Errorf("exts %v, err %v", exts, err)
            }
            name := link.Subreddit + "-" + link.Title + exts[0]
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

