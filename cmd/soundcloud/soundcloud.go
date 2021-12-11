package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/soundcloud"
   "net/http"
   "os"
   "strings"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("soundcloud [-i] [URL]")
      flag.PrintDefaults()
      return
   }
   addr := flag.Arg(0)
   track, err := soundcloud.Resolve(addr)
   if err != nil {
      panic(err)
   }
   if info {
      fmt.Printf("%+v\n", track)
      return
   }
   pro, err := track.Progressive()
   if err != nil {
      panic(err)
   }
   fmt.Println("GET", pro.URL)
   res, err := http.Get(pro.URL)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   name := track.User.Username + "-" + track.Title + ".mp3"
   file, err := os.Create(strings.Map(mech.Clean, name))
   if err != nil {
      panic(err)
   }
   defer file.Close()
   if _, err := file.ReadFrom(res.Body); err != nil {
      panic(err)
   }
}
