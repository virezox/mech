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
   media, err := cwtv.Media(play)
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
   sort.Slice(mas.Stream, func(a, b int) bool {
      return mas.Stream[a].Bandwidth < mas.Stream[b].Bandwidth
   })
   if info {
      done := make(map[hls.Stream]bool)
      for _, str := range mas.Stream {
         str.URI = ""
         if !done[str] {
            done[str] = true
            fmt.Println(str)
         }
      }
   } else {
      uris := mas.URIs(func(str hls.Stream) bool {
         return str.Bandwidth >= bandwidth
      })
      for _, uri := range uris {
         fmt.Println(uri)
         /*
         vid, err := cwtv.NewVideo(guid)
         if err != nil {
            return err
         }
         if err := download(stream, vid.Name()); err != nil {
            return err
         }
         */
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
