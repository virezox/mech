package main

import (
   "fmt"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "github.com/89z/mech/nbc"
   "io"
)

func new_master(guid int64, bandwidth int, info bool) error {
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
   stream := master.Stream.Reduce(hls.Bandwidth(bandwidth))
   if info {
      for _, item := range master.Stream {
         if item.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
   } else {
      return download(stream.URI, page.Analytics.ConvivaAssetName)
   }
   return nil
}

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
   pro := os.Progress_Chunks(file, len(seg.URI))
   for _, addr := range seg.URI {
      res, err := nbc.Client.Level(0).Redirect(nil).Get(addr)
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
