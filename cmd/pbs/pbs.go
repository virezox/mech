package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/pbs"
   "io"
   "net/http"
   "net/url"
   "os"
)

func doWidget(address, audio string, video int64, info bool) error {
   getter, err := pbs.NewWidgeter(address)
   if err != nil {
      return err
   }
   widget, err := getter.Widget()
   if err != nil {
      return err
   }
   addr := widget.HLS()
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
      fmt.Println(widget)
      for _, stream := range master.Streams {
         fmt.Println(stream)
      }
      for _, media := range master.Media {
         fmt.Println(media)
      }
   } else {
      media := master.Audio(audio)
      if media != nil {
         err := download(media.URI, widget.Slug + hls.AAC)
         if err != nil {
            return err
         }
      }
      stream := master.Stream(video)
      return download(stream.URI, widget.Slug + hls.TS)
   }
   return nil
}

func download(addr *url.URL, name string) error {
   fmt.Println("GET", addr)
   res, err := http.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.NewScanner(res.Body).Segment(res.Request.URL)
   if err != nil {
      return err
   }
   file, err := os.Create(name)
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
