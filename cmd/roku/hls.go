package main

import (
   "fmt"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "github.com/89z/mech/roku"
   "io"
   "net/url"
)

func (d downloader) HLS(bandwidth int64) error {
   video, err := d.Content.HLS()
   if err != nil {
      return err
   }
   res, err := roku.Client.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   stream := master.Streams.Get_Bandwidth(bandwidth)
   if !d.info {
      addr, err := res.Request.URL.Parse(stream.URI)
      if err != nil {
         return err
      }
      return download_HLS(addr, d.Base())
   }
   return nil
}
