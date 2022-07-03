package main

import (
   "fmt"
   "github.com/89z/mech/paramount"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "io"
)

func (d downloader) HLS(bandwidth int64) error {
   addr, err := paramount.New_Media(d.guid).HLS()
   if err != nil {
      return err
   }
   res, err := paramount.Client.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   stream := master.Streams.Get_Bandwidth(bandwidth)
   if d.info {
      fmt.Println(d.Title)
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream.URI, d.Base())
   }
   return nil
}
