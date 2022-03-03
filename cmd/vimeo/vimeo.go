package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/vimeo"
   "net/http"
   "net/url"
   "os"
   "path"
   "sort"
)

func doClip(clip *vimeo.Clip, addr, name string, info bool) error {
   web, err := vimeo.NewJsonWeb()
   if err != nil {
      return err
   }
   if addr != "" {
      clip, err = vimeo.NewClip(addr)
      if err != nil {
         return err
      }
   }
   video, err := web.Video(clip)
   if err != nil {
      return err
   }
   sort.Slice(video.Download, func(a, b int) bool {
      return video.Download[a].Height < video.Download[b].Height
   })
   if info {
      fmt.Println(video)
   } else {
      for _, down := range video.Download {
         if down.Public_Name == name {
            err := download(down)
            if err != nil {
               return err
            }
         }
      }
   }
   return nil
}

func download(down vimeo.Download) error {
   fmt.Println("GET", down.Link)
   res, err := http.Get(down.Link)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   addr, err := url.Parse(down.Link)
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
