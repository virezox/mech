package main

import (
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
   "os"
   "path/filepath"
)

func login(username, password string) error {
   login, err := instagram.NewLogin(username, password)
   if err != nil {
      return err
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cache = filepath.Join(cache, "mech")
   os.Mkdir(cache, os.ModePerm)
   cache = filepath.Join(cache, "instagram.json")
   fmt.Println("Create", cache)
   return login.Create(cache)
}

func mediaItems(shortcode string, auth bool) ([]instagram.MediaItem, error) {
   if auth {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      log, err := instagram.OpenLogin(cache + "/mech/instagram.json")
      if err != nil {
         return nil, err
      }
      return log.MediaItems(shortcode)
   }
   return instagram.MediaItems(shortcode)
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
   file, err := os.Create(filepath.Base(addr.Path))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
