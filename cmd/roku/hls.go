package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "io"
   "net/http"
   "net/url"
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
   master, err := hls.NewScanner(res.Body).Master()
   if err != nil {
      return err
   }
   stream := master.Streams.GetBandwidth(bandwidth)
   if !d.info {
      addr, err := stream.URI(res.Request.URL)
      if err != nil {
         return err
      }
      return downloadHLS(addr, d.Base())
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

func downloadHLS(addr *url.URL, base string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.NewScanner(res.Body).Segment()
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
      addr, err := info.URI(res.Request.URL)
      if err != nil {
         return err
      }
      res, err := http.Get(addr.String())
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
