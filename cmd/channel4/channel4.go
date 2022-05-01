package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/dash"
   "github.com/89z/mech/channel4"
   "io"
   "net/http"
   "net/url"
   "os"
   "sort"
)

func download(temp dash.Template, base *url.URL) error {
   addrs, err := temp.URLs(base)
   if err != nil {
      return err
   }
   file, err := os.Create(temp.Initialization)
   if err != nil {
      return err
   }
   pro := format.ProgressChunks(file, len(addrs))
   for _, addr := range addrs {
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
   return file.Close()
}

func doManifest(id, address string, bandwidth int64, info bool) error {
   if id == "" {
      id = channel4.ProgramID(address)
   }
   stream, err := channel4.NewStream(id)
   if err != nil {
      return err
   }
   widevine := stream.Widevine()
   fmt.Println("GET", widevine)
   res, err := http.Get(widevine)
   if err != nil {
      return err
   }
   period, err := dash.NewPeriod(res.Body)
   if err != nil {
      return err
   }
   video := period.Video()
   if bandwidth >= 1 {
      sort.Sort(dash.Bandwidth{video, bandwidth})
   }
   for _, rep := range video.Representation {
      if info {
         fmt.Println(rep)
      } else {
         temp := video.SegmentTemplate.Replace(rep)
         return download(temp, res.Request.URL)
      }
   }
   return res.Body.Close()
}
