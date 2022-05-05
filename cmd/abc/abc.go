package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/abc"
   "net/http"
   "os"
)

func newMaster(addr string, bandwidth int, info bool) error {
   route, err := abc.NewRoute(addr)
   if err != nil {
      return err
   }
   video, err := route.Video()
   if err != nil {
      return err
   }
   if err := video.Authorize(); err != nil {
      return err
   }
   addr = video.ULNK()
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(video)
      for _, stream := range master.Streams {
         fmt.Println(stream)
      }
   } else {
      stream := master.Stream(bandwidth)
      return download(stream, video.Base())
   }
   return nil
}

func newSegment(stream *hls.Stream) (*hls.Segment, error) {
   fmt.Println("GET", stream.URI)
   res, err := http.Get(stream.URI.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewScanner(res.Body).Segment(res.Request.URL)
}

func download(stream *hls.Stream, base string) error {
   seg, err := newSegment(stream)
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
