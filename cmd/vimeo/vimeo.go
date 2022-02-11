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

func newClip(clipID, unlistedHash int64, addr string) (*vimeo.Clip, error) {
   if clipID >= 1 {
      return &vimeo.Clip{clipID, unlistedHash}, nil
   }
   return vimeo.NewClip(addr)
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
