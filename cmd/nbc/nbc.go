package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/nbc"
   "net/http"
   "os"
)

func newMaster(addr string) (*hls.Master, error) {
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewMaster(res.Request.URL, res.Body)
}

func doManifest(guid, form int64, info bool) error {
   vod, err := nbc.NewAccessVOD(guid)
   if err != nil {
      return err
   }
   vid, err := nbc.NewVideo(guid)
   if err != nil {
      return err
   }
   file, err := os.Create(vid.Name())
   if err != nil {
      return err
   }
   defer file.Close()
   mas, err := newMaster(vod.ManifestPath)
   if err != nil {
      return err
   }
   res, err := http.Get(mas.Stream[0].URI)
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
