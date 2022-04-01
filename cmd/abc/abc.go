package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/abc"
   "net/http"
   "os"
   "sort"
)

func doManifest(addr string, bandwidth int, info bool) error {
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
   master, err := hls.NewMaster(res.Request.URL, res.Body)
   if err != nil {
      return err
   }
   sort.Sort(hls.Bandwidth{master, bandwidth})
   if info {
      fmt.Println(video)
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

func newSegment(stream hls.Stream) (*hls.Segment, error) {
   fmt.Println("GET", stream.URI)
   res, err := http.Get(stream.URI.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return hls.NewSegment(res.Request.URL, res.Body)
}

func download(stream hls.Stream, video *abc.Video) error {
   seg, err := newSegment(stream)
   if err != nil {
      return err
   }
   fmt.Println("GET", seg.Key.URI)
   res, err := http.Get(seg.Key.URI.String())
   if err != nil {
      return err
   }
   block, err := hls.NewCipher(res.Body)
   if err != nil {
      return err
   }
   file, err := os.Create(video.Base() + seg.Ext())
   if err != nil {
      return err
   }
   for i, info := range seg.Info {
      fmt.Print(seg.Progress(i))
      res, err := http.Get(info.URI.String())
      if err != nil {
         return err
      }
      buf, err := block.Decrypt(info, res.Body)
      if err != nil {
         return err
      }
      if _, err := file.Write(buf); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   if err := res.Body.Close(); err != nil {
      return err
   }
   return file.Close()
}
