package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/tiktok"
   "net/http"
   "os"
   "strings"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("tiktok [flags] [URL]")
      flag.PrintDefaults()
      return
   }
   addr := flag.Arg(0)
   item, err := tiktok.NewItemStruct(addr)
   if err != nil {
      panic(err)
   }
   req, err := item.Request()
   if err != nil {
      panic(err)
   }
   if info {
      mech.LogLevel.Dump(2, req)
   } else {
      err := get(req, item)
      if err != nil {
         panic(err)
      }
   }
}

func get(req *http.Request, item *tiktok.ItemStruct) error {
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   ext, err := mech.ExtensionByType(res.Header.Get("Content-Type"))
   if err != nil {
      return err
   }
   name := item.Author.UniqueID + "-" + item.ID + ext
   file, err := os.Create(strings.Map(mech.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      return err
   }
   return nil
}
