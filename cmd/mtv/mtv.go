package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/mtv"
   "net/http"
   "os"
   "sort"
)

func download(stream hls.Stream, name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   fmt.Println(stream.URI)
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

func doManifest(addr string, bandwidth int64, info bool) error {
   prop, err := mtv.NewItem(addr).Property()
   if err != nil {
      return err
   }
   top, err := prop.Topaz()
   if err != nil {
      return err
   }
   fmt.Println(top.StitchedStream.Source)
   res, err := http.Get(top.StitchedStream.Source)
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
   for _, stream := range mas.Stream {
      if info {
         stream.URI = ""
         fmt.Println(stream)
      } else if stream.Bandwidth >= bandwidth {
         err := download(stream, "ignore.mp4")
         if err != nil {
            return err
         }
         break
      }
   }
   return nil
}


