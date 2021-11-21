package main

import (
   "fmt"
   "github.com/89z/mech/reddit"
   "github.com/89z/parse/m3u"
   "net/http"
   "os"
   "path"
   "strconv"
)

func (c choice) HLS(link *reddit.Link) error {
   hlss, err := link.HLS()
   if err != nil {
      return err
   }
   for _, hls := range hlss {
      if c.format {
         fmt.Printf("%+v\n", hls)
      } else if c.ids[strconv.Itoa(hls.ID)] {
         fmt.Println("GET", hls.URI)
         res, err := http.Get(hls.URI)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         prefix, _ := path.Split(hls.URI)
         for key := range m3u.NewByteRange(res.Body) {
            fmt.Println("GET", prefix + key)
            res, err := http.Get(prefix + key)
            if err != nil {
               return err
            }
            defer res.Body.Close()
            file, err := os.Create(key)
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
