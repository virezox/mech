package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bbc"
   "net/http"
   "os"
   "time"
)

func media(item *bbc.NewsItem, info bool, form int64) error {
   media, err := item.Media()
   if err != nil {
      return err
   }
   streams, err := media.Streams()
   if err != nil {
      return err
   }
   for _, stream := range streams {
      if info {
         fmt.Println(stream)
      } else if stream.ID == form {
         name, err := media.Name(item)
         if err != nil {
            return err
         }
         if err := download(name, stream); err != nil {
            return err
         }
      }
   }
   return nil
}

func download(name string, stream bbc.Stream) error {
   file, err := os.Create(name)
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
      res, err := http.Get(info)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      fmt.Print(format.PercentInt(i, len(infos)))
      fmt.Print("\t")
      fmt.Print(format.Size.Get(size))
      fmt.Print("\t")
      fmt.Println(format.Rate.Get(size/time.Since(begin).Seconds()))
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
      size += float64(res.ContentLength)
   }
   return nil
}
