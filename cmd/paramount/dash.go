package main

import (
   "fmt"
   "github.com/89z/mech/paramount"
)

func doDASH(guid string, bandwidth int64, info bool) error {
   media, err := paramount.DASH(guid)
   if err != nil {
      return err
   }
   video, err := media.Video()
   if err != nil {
      return err
   }
   fmt.Println("GET", video.Src)
   return nil
}
