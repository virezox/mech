package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/apple"
   "net/http"
   "os"
   "path"
   "strings"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("apple [-i] [URL]")
      flag.PrintDefaults()
      return
   }
   addr := flag.Arg(0)
   mech.Verbose(true)
   audio, err := apple.NewAudio(addr)
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Printf("%+v\n", audio)
      return
   }
   for _, asset := range audio.D {
      err := download(asset.Attributes)
      if err != nil {
         panic(err)
      }
   }
}

func download(attr apple.Attributes) error {
   fmt.Println("GET", attr.AssetURL)
   res, err := http.Get(attr.AssetURL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := attr.ArtistName + "-" + attr.Name + path.Ext(attr.AssetURL)
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
