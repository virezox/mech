package main

import (
   "github.com/89z/format"
   "github.com/89z/mech/nbc"
   "net/http"
   "os"
   "strings"
   "time"
)

func video(guid uint64, info bool) (*nbc.Video, error) {
   if info {
      return nil, nil
   }
   return nbc.NewVideo(guid)
}

func download(vid *nbc.Video, stream nbc.Stream) error {
   name := vid.Name() + "-" + stream.Resolution + ".mp4"
   file, err := os.Create(strings.Map(format.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   infos, err := stream.Information()
   if err != nil {
      return err
   }
   begin := time.Now()
   var size float64
   for i, info := range infos {
      res, err := http.Get(info.URI)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      os.Stdout.WriteString(format.PercentInt(i, len(infos)))
      os.Stdout.WriteString("\t")
      os.Stdout.WriteString(format.Size.Get(size))
      os.Stdout.WriteString("\t")
      os.Stdout.WriteString(format.Rate.Get(size/time.Since(begin).Seconds()))
      os.Stdout.WriteString("\n")
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      size += float64(res.ContentLength)
   }
   return nil
}
