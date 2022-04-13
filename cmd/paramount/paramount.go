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

func doManifest(guid string, bandwidth int, info bool) error {
   media, err := paramount.NewMedia(guid)
   if err != nil {
      return err
   }
   video, err := media.Video()
   if err != nil {
      return err
   }
   res, err := http.Get(video.Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   paramount.LogLevel.Dump(res.Request)
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(video.Title)
   }
   if bandwidth >= 1 {
      sort.Sort(hls.Bandwidth{master, bandwidth})
   }
   for _, stream := range master.Stream {
      if info {
         fmt.Println(stream)
      } else {
         return download(stream, video)
      }
   }
   return nil
}

func download(stream hls.Stream, video *paramount.Video) error {
   seg, err := newSegment(stream.URI.String())
   if err != nil {
      return err
   }
   res, err := http.Get(seg.Key.String())
   if err != nil {
      return err
   }
   paramount.LogLevel.Dump(res.Request)
   block, err := hls.NewCipher(res.Body)
   if err != nil {
      return err
   }
   if err := res.Body.Close(); err != nil {
      return err
   }
   file, err := os.Create(video.Base() + seg.Ext())
   if err != nil {
      return err
   }
   pro := format.NewProgress(file, len(seg.Info))
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
   return file.Close()
}

func newSegment(addr string) (*hls.Segment, error) {
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   paramount.LogLevel.Dump(res.Request)
   return hls.NewScanner(res.Body).Segment(res.Request.URL)
}
