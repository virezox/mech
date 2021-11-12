package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/pbs"
   "net/http"
   "os"
   "strings"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("pbs [-i] [URL]")
      flag.PrintDefaults()
      return
   }
   addr := flag.Arg(0)
   mech.Verbose = true
   slug, err := pbs.Slug(addr)
   if err != nil {
      panic(err)
   }
   asset, err := pbs.NewAsset(slug)
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Printf("%+v\n", asset)
      return
   }
   for _, video := range asset.Resource.MP4_Videos {
      err := download(asset.Resource.Title, video)
      if err != nil {
         panic(err)
      }
   }
}

func download(title string, video pbs.Video) error {
   fmt.Println("GET", video.URL)
   res, err := http.Get(video.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := title + "-" + video.Profile + ".mp4"
   file, err := os.Create(strings.Map(mech.Clean, name))
   if err != nil {
      return err
   }
   pro := pbs.NewProgress(res)
   file.ReadFrom(pro)
   return file.Close()
}
