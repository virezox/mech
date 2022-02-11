package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/tiktok"
   "net/http"
   "os"
   "strings"
)

func get(det *tiktok.AwemeDetail, output string) error {
   addr, err := det.URL()
   if err != nil {
      return err
   }
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if output == "" {
      ext, err := format.ExtensionByType(res.Header.Get("Content-Type"))
      if err != nil {
         return err
      }
      name := det.Author.Unique_ID + "-" + det.Aweme_ID + ext
      output = strings.Map(format.Clean, name)
   }
   file, err := os.Create(output)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
