package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/roku"
   "io"
   "net/http"
   "os"
)

func doManifest(guid int64, bandwidth int, info bool) error {
   vod, err := roku.NewAccessVOD(guid)
   if err != nil {
      return err
   }
   fmt.Println("GET", vod.ManifestPath)
   res, err := http.Get(vod.ManifestPath)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   if info {
      for _, stream := range master.Streams {
         fmt.Println(stream)
      }
   } else {
      stream := master.Stream(bandwidth)
      video, err := roku.NewVideo(guid)
      if err != nil {
         return err
      }
      return download(stream, video)
   }
   return nil
}

func download(stream *hls.Stream, video *roku.Video) error {
   fmt.Println("GET", stream.URI)
   res, err := http.Get(stream.URI.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.NewScanner(res.Body).Segment(res.Request.URL)
   if err != nil {
      return err
   }
   file, err := os.Create(video.Base() + hls.TS)
   if err != nil {
      return err
   }
   defer file.Close()
   pro := format.ProgressChunks(file, len(seg.Info))
   for _, info := range seg.Info {
      res, err := http.Get(info.URI.String())
      if err != nil {
         return err
      }
      pro.AddChunk(res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
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
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      roku.LogLevel = 1
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
