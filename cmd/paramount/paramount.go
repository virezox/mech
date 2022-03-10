package main

import (
   "flag"
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/paramount"
   "net/http"
   "os"
   "sort"
)

func doManifest(guid, bandwidth int64, info bool) error {
   vod, err := paramount.NewAccessVOD(guid)
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
      vid, err := paramount.NewVideo(guid)
      if err != nil {
         return err
      }
      if err := download(str, vid); err != nil {
         return err
      }
   }
   return nil
}

func download(str *hls.Stream, video *paramount.Video) error {
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

func main() {
   // b
   var guid int64
   flag.Int64Var(&guid, "b", 0, "GUID")
   // f
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 4_000_000, "min bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      paramount.LogLevel = 1
   }
   if guid >= 1 {
      err := doManifest(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
