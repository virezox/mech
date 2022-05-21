package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "io"
   "net/http"
   "os"
)

func (d downloader) HLS(bandwidth int64) error {
   video, err := d.Content.HLS()
   if err != nil {
      return err
   }
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   stream := master.Stream(bandwidth)
   if !d.info {
      return downloadHLS(stream, d.Base())
   }
   fmt.Println(d.Content)
   for _, each := range master.Streams {
      if each.Bandwidth == stream.Bandwidth {
         fmt.Print("!")
      }
      fmt.Println(each)
   }
   return nil
}

func downloadHLS(stream *hls.Stream, base string) error {
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
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
