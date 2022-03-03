package main

import (
   "flag"
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/mtv"
   "net/http"
   "os"
)

func doManifest(guid, bandwidth int64, info bool) error {
   vod, err := mtv.NewAccessVOD(guid)
   if err != nil {
      return err
   }
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
         vid, err := mtv.NewVideo(guid)
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

func main() {
   // b
   var guid int64
   flag.Int64Var(&guid, "b", 0, "GUID")
   // f
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 5480000, "bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      mtv.LogLevel = 1
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
