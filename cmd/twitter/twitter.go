package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/twitter"
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
      fmt.Println("twitter [-i] [URL]")
      flag.PrintDefaults()
      return
   }
   addr := flag.Arg(0)
   audio, err := twitter.NewAudio(addr)
   if err != nil {
      panic(err)
   }
   for _, asset := range audio.D {
      if info {
         fmt.Printf("%+v\n", asset.Attributes)
      } else {
         err := download(asset.Attributes)
         if err != nil {
            panic(err)
         }
      }
   }
}

func download(attr twitter.Attributes) error {
   addr := attr.AssetURL.String()
   fmt.Println("GET", addr)
   res, err := http.Get(addr)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := attr.ArtistName + "-" + attr.Name + path.Ext(addr)
   file, err := os.Create(strings.Map(mech.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   pro := mech.NewProgress(res)
   if _, err := file.ReadFrom(pro); err != nil {
      return err
   }
   return nil
}
