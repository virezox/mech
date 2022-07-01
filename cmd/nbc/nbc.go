package main

import (
   "fmt"
   "github.com/89z/std/hls"
   "github.com/89z/std/os"
   "github.com/89z/mech/nbc"
   "io"
   "sort"
)

func download(addr, base string) error {
   res, err := nbc.Client.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.New_Scanner(res.Body).Segment()
   if err != nil {
      return err
   }
   file, err := os.Create(base + ".ts")
   if err != nil {
      return err
   }
   defer file.Close()
   pro := os.Progress_Chunks(file, len(seg.Clear))
   for _, clear := range seg.Clear {
      res, err := nbc.Client.Level(0).Redirect(nil).Get(clear)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
func new_master(guid, bandwidth int64, info bool) error {
   page, err := nbc.New_Bonanza_Page(guid)
   if err != nil {
      return err
   }
   video, err := page.Video()
   if err != nil {
      return err
   }
   res, err := nbc.Client.Get(video.ManifestPath)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.New_Scanner(res.Body).Master()
   if err != nil {
      return err
   }
   sort.Slice(master.Streams, func(a, b int) bool {
      return master.Streams[a].Bandwidth < master.Streams[b].Bandwidth
   })
   stream := master.Streams.Get_Bandwidth(bandwidth)
   if info {
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream.URI, page.Analytics.ConvivaAssetName)
   }
   return nil
}
