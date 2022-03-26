package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/vimeo"
   "net/http"
   "net/url"
   "os"
   "path"
)

func doAuth(clip *vimeo.Clip, password string, height int, info bool) error {
   check, err := clip.Check(password)
   if err != nil {
      return err
   }
   for _, prog := range check.Request.Files.Progressive {
      if info {
         fmt.Println(prog)
      } else if prog.Height == height {
         return download(prog.URL)
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
   pro := format.NewProgress(res)
   if _, err := file.ReadFrom(pro); err != nil {
      return err
   }
   return nil
}

func doAnon(clip *vimeo.Clip, height int, info bool) error {
   web, err := vimeo.NewJsonWeb()
   if err != nil {
      return err
   }
   video, err := web.Video(clip)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(video)
   } else {
      for _, down := range video.Download {
         if down.Height == height {
            return download(down.Link)
         }
      }
   }
   return nil
}
