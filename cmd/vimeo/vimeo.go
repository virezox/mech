package main

import (
   "fmt"
   "github.com/89z/mech/vimeo"
   "github.com/89z/rosso/os"
   "io"
   "net/http"
   "net/url"
   "path"
)

func (f flags) vimeo() error {
   web, err := vimeo.New_JSON_Web()
   if err != nil {
      return err
   }
   clip, err := vimeo.New_Clip(f.address)
   if err != nil {
      panic(err)
   }
   video, err := web.Video(clip)
   if err != nil {
      return err
   }
   if f.info {
      fmt.Println(video)
   } else {
      for _, down := range video.Download {
         if down.Height == f.height {
            return download(down.Link)
         }
      }
   }
   return nil
}

func (f flags) vhx() error {
   embed, err := vimeo.New_Embed(f.address)
   if err != nil {
      return err
   }
   config, err := embed.Config()
   if err != nil {
      return err
   }
   if f.info {
      fmt.Println(config)
   } else {
      for _, pro := range config.Request.Files.Progressive {
         if pro.Height == f.height {
            return download(pro.URL)
         }
      }
   }
   return nil
}

func download(address string) error {
   fmt.Println("GET", address)
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   addr, err := url.Parse(address)
   if err != nil {
      return err
   }
   file, err := os.Create(path.Base(addr.Path))
   if err != nil {
      return err
   }
   defer file.Close()
   pro := os.Progress_Bytes(file, res.ContentLength)
   if _, err := io.Copy(pro, res.Body); err != nil {
      return err
   }
   return nil
}
