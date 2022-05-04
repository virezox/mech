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
   "sort"
)

func doWidget(address string, bandwidth int, info bool) error {
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
   fmt.Println(widget)
   sort.Sort(hls.Bandwidth{master, bandwidth})
   for _, video := range master.Stream {
      if info {
         fmt.Println(video)
      } else {
         audio := master.Audio(video)
         if audio != nil {
            err := download(audio.URI, widget.Slug + hls.AAC)
            if err != nil {
               return err
            }
         }
         return download(video.URI, widget.Slug + hls.TS)
      }
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
