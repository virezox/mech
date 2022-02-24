package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/tiktok"
   "net/http"
   "os"
   "strings"
)

func get(det *tiktok.AwemeDetail) error {
   addr := det.URL()
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   ext, err := format.ExtensionByType(res.Header.Get("Content-Type"))
   if err != nil {
      return err
   }
   name := det.Author.Unique_ID + "-" + det.Aweme_ID + ext
   file, err := os.Create(strings.Map(format.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
