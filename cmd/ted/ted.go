package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/ted"
   "net/http"
   "os"
   "path"
)

func process(slug string, info bool, bitrate int64) error {
   talk, err := ted.NewTalkResponse(slug)
   if err != nil {
      return err
   }
   for _, vid := range talk.Downloads.Video {
      if info {
         fmt.Println(vid)
      } else if vid.Bitrate == bitrate {
         addr := vid.GetURL()
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         file, err := os.Create(path.Base(addr))
         if err != nil {
            return err
         }
         defer file.Close()
         pro := format.NewProgress(res)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
      }
   }
   return nil
}
