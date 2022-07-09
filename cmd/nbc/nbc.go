package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/nbc"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
)

func main() {
   // b
   var guid int64
   flag.Int64Var(&guid, "b", 0, "GUID")
   // f
   var bandwidth int
   // nbc.com/saturday-night-live/video/march-12-zoe-kravitz/9000199371
   // nbc.com/saturday-night-live/video/may-15-keeganmichael-key/4358937
   flag.IntVar(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      nbc.Client.Log_Level = 2
   }
   if guid >= 1 {
      err := new_master(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
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
   distance := hls.Bandwidth(bandwidth)
   stream := master.Stream.Reduce(func(carry, item hls.Stream) bool {
      return distance(item) < distance(carry)
   })
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
