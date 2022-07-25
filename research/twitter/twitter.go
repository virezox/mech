package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/research/twitter"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
   "net/http"
)

type flags struct {
   bitrate int64
   mech.Stream
   space string
}

func main() {
   var f flags
   // c
   flag.StringVar(&f.space, "b", "", "space ID")
   // f
   flag.Int64Var(&f.bitrate, "f", 1, "status bitrate")
   // i
   flag.BoolVar(&f.Info, "i", false, "info")
   flag.Parse()
   if f.space != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) download() error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   space, err := guest.AudioSpace(f.space)
   if err != nil {
      return err
   }
   if f.Info {
      fmt.Println(space)
   } else {
      source, err := guest.Source(space)
      if err != nil {
         return err
      }
      fmt.Println("GET", source.Location)
      res, err := http.Get(source.Location)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      seg, err := hls.New_Scanner(res.Body).Segment()
      if err != nil {
         return err
      }
      file, err := os.Create(space.Base() + ".aac")
      if err != nil {
         return err
      }
      defer file.Close()
      pro := os.Progress_Chunks(file, len(seg.URI))
      for _, ref := range seg.URI {
         req, err := http.NewRequest("GET", ref, nil)
         if err != nil {
            return err
         }
         req.URL = res.Request.URL.ResolveReference(req.URL)
         res, err := twitter.Client.Level(0).Do(req)
         if err != nil {
            return err
         }
         pro.Add_Chunk(res.ContentLength)
         if _, err := io.Copy(pro, res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
