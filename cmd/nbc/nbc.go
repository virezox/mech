package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/nbc"
   "io"
   "net/http"
   "os"
   "sort"
)

func doManifest(guid int64, bandwidth int, info bool) error {
   vod, err := nbc.NewAccessVOD(guid)
   if err != nil {
      return err
   }
   fmt.Println("GET", vod.ManifestPath)
   res, err := http.Get(vod.ManifestPath)
   if err != nil {
      return err
   }
   master, err := hls.NewScanner(res.Body).Master(res.Request.URL)
   if err != nil {
      return err
   }
   if bandwidth >= 1 {
      sort.Sort(hls.Bandwidth{master, bandwidth})
   }
   for _, stream := range master.Stream {
      if info {
         fmt.Println(stream)
      } else {
         video, err := nbc.NewVideo(guid)
         if err != nil {
            return err
         }
         return download(stream, video)
      }
   }
   return res.Body.Close()
}

func download(stream hls.Stream, video *nbc.Video) error {
   fmt.Println("GET", stream.URI)
   res, err := http.Get(stream.URI.String())
   if err != nil {
      return err
   }
   seg, err := hls.NewScanner(res.Body).Segment(res.Request.URL)
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
   pro := format.ProgressChunks(file, len(seg.Info))
   for _, info := range seg.Info {
      res, err := http.Get(info.URI.String())
      if err != nil {
         return err
      }
      if res.StatusCode != http.StatusOK {
         return errorString(res.Status)
      }
      pro.AddChunk(res.ContentLength)
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return file.Close()
}

type errorString string

func (e errorString) Error() string {
   return string(e)
}
