package main

import (
   "fmt"
   "github.com/89z/mech/instagram"
   "net/http"
   "net/url"
   "os"
   "path"
   "time"
)

func graphqlPath(shortcode string, info bool) error {
   graph, err := instagram.NewGraphQL(shortcode)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(graph)
   } else {
      for _, addr := range graph.URLs() {
         err := download(addr)
         if err != nil {
            return err
         }
         time.Sleep(99 * time.Millisecond)
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
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}

func mediaItemPath(shortcode string, info bool) error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   login, err := instagram.OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      return err
   }
   id, err := instagram.GetID(shortcode)
   if err != nil {
      return err
   }
   items, err := login.MediaItems(id)
   if err != nil {
      return err
   }
   for _, item := range items {
      form, err := item.Format()
      if err != nil {
         return err
      }
      fmt.Println(form)
   }
   return nil
}

