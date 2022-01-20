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
   var size int64
   for i, info := range infos {
      res, err := http.Get(info.URI)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      format.PercentInt(os.Stdout, i, len(infos))
      os.Stdout.WriteString("\t")
      format.Size.Int64(os.Stdout, size)
      os.Stdout.WriteString("\t")
      format.Rate.Int64(os.Stdout, size/time.Since(begin).Milliseconds()*1000)
      os.Stdout.WriteString("\n")
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      size += res.ContentLength
   }
   return nil
}
