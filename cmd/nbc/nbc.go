package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/nbc"
   "net/http"
   "os"
   "sort"
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
   sort.Slice(mas.Stream, func(a, b int) bool {
      return mas.Stream[a].Bandwidth < mas.Stream[b].Bandwidth
   })
   if info {
      for _, str := range mas.Stream {
         str.URI = nil
         fmt.Println(str)
      }
   } else {
      str := mas.GetStream(func (s hls.Stream) bool {
         return s.Bandwidth >= bandwidth
      })
      vid, err := nbc.NewVideo(guid)
      if err != nil {
         return err
      }
      if err := download(str, vid); err != nil {
         return err
      }
   }
   return nil
}

func download(str *hls.Stream, video *nbc.Video) error {
   fmt.Println("GET", str.URI)
   res, err := http.Get(str.URI.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   str.URI = nil
   fmt.Println(str)
   seg, err := hls.NewSegment(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(video.Base() + seg.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   for i, info := range seg.Info {
      fmt.Print(seg.Progress(i))
      res, err := http.Get(info.URI.String())
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
