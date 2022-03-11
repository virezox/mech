package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/cwtv"
   "net/http"
   "os"
   "sort"
)

func doManifest(addr string, bandwidth int64, info bool) error {
   play, err := cwtv.GetPlay(addr)
   if err != nil {
      return err
   }
   video, err := cwtv.NewVideo(play)
   if err != nil {
      return err
   }
   media, err := video.Media()
   if err != nil {
      return err
   }
   fmt.Println("GET", media)
   res, err := http.Get(media.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   mas, err := hls.NewMaster(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   if info {
      done := make(map[hls.Stream]bool)
      for _, str := range mas.Stream {
         str.URI = nil
         if !done[str] {
            done[str] = true
            fmt.Println(str)
         }
      }
   } else {
      sort.Sort(hls.Bandwidth{mas, bandwidth})
      err := download(video, mas.Stream[0])
      if err != nil {
         return err
      }
   }
   return nil
}

func download(video *cwtv.Video, str hls.Stream) error {
   fmt.Println("GET", str.URI)
   res, err := http.Get(str.URI.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.NewSegment(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(video.Base() + seg.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   str.URI = nil
   fmt.Println(str)
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
