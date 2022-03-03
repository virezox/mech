package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/nbc"
   "net/http"
   "os"
)

func doManifest(guid, bandwidth int64, info bool) error {
   vod, err := nbc.NewAccessVOD(guid)
   if err != nil {
      return err
   }
   fmt.Println("GET", vod.ManifestPath)
   res, err := http.Get(vod.ManifestPath)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   mas, err := hls.NewMaster(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   for _, stream := range mas.Stream {
      if info {
         stream.URI = ""
         fmt.Println(stream)
      } else if stream.Bandwidth == bandwidth {
         vid, err := nbc.NewVideo(guid)
         if err != nil {
            return err
         }
         if err := download(stream, vid.Name()); err != nil {
            return err
         }
      }
   }
   return nil
}

func download(stream hls.Stream, name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   fmt.Println("GET", stream.URI)
   res, err := http.Get(stream.URI)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.NewSegment(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   for i, info := range seg.Info {
      fmt.Println(i, len(seg.Info))
      res, err := http.Get(info.URI)
      if err != nil {
         return err
      }
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
