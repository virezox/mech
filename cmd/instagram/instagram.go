package main

import (
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
   "os"
   "path"
)

func newMedia(shortcode string, auth bool) (*instagram.Media, error) {
   if auth {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      log, err := instagram.OpenLogin(cache + "/mech/instagram.json")
      if err != nil {
         return nil, err
      }
      return log.Media(shortcode)
   }
   return instagram.NewMedia(shortcode)
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
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
