package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/paramount"
   "net/http"
   "os"
   "sort"
   "time"
)

func doManifest(guid, address string, bandwidth int, info bool) error {
   if guid == "" {
      guid = paramount.GUID(address)
   }
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

func get(addr string) (*http.Response, error) {
   req, err := http.NewRequest("", addr, nil)
   if err != nil {
      return nil, err
   }
   // 9 seconds is too long
   tr := &http.Transport{IdleConnTimeout: 8*time.Second}
   res, err := tr.RoundTrip(req)
   if err != nil {
      fmt.Println("RETRY")
      return tr.RoundTrip(req)
   }
   return res, nil
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
   defer file.Close()
   pro := format.ProgressChunks(file, len(seg.Info))
   for _, info := range seg.Info {
      res, err := get(info.URI.String())
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
   res, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   paramount.LogLevel.Dump(res.Request)
   return hls.NewScanner(res.Body).Segment(res.Request.URL)
}
