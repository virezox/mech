package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/paramount"
   "net/http"
   "os"
   "sort"
)

func doHLS(guid string, bandwidth int64, info bool) error {
   media, err := paramount.HLS(guid)
   if err != nil {
      return err
   }
   video, err := media.Video()
   if err != nil {
      return err
   }
   fmt.Println("GET", video.Src)
   res, err := http.Get(video.Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   sort.Slice(master.Streams, func(a, b int) bool {
      return master.Streams[a].Bandwidth < master.Streams[b].Bandwidth
   })
   stream := master.Stream(bandwidth)
   if info {
      fmt.Println(video.Title)
      for _, each := range master.Streams {
         if each.Bandwidth == stream.Bandwidth {
            fmt.Print("!")
         }
         fmt.Println(each)
      }
   } else {
      return download(stream, video.Base())
   }
   return nil
}

func download(stream *hls.Stream, base string) error {
   seg, err := newSegment(stream.URI.String())
   if err != nil {
      return err
   }
   fmt.Println("GET", seg.Key)
   res, err := http.Get(seg.Key.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   block, err := hls.NewCipher(res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(base + hls.TS)
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
      if _, err := block.Copy(pro, res.Body, info.IV); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

func newSegment(addr string) (*hls.Segment, error) {
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewScanner(res.Body).Segment(res.Request.URL)
}
