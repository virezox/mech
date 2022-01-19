package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/soundcloud"
   "net/http"
   "os"
   "strings"
)

func main() {
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if flag.NArg() == 1 {
      addr := flag.Arg(0)
      track, err := soundcloud.Resolve(addr)
      if err != nil {
         panic(err)
      }
      if info {
         fmt.Printf("%+v\n", track)
      } else {
         err := download(track)
         if err != nil {
            panic(err)
         }
      }
   } else {
      fmt.Println("soundcloud [-i] [URL]")
      flag.PrintDefaults()
   }
}

func download(track *soundcloud.Track) error {
   pro, err := track.Progressive()
   if err != nil {
      return err
   }
   fmt.Println("GET", pro.URL)
   res, err := http.Get(pro.URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   name := track.User.Username + "-" + track.Title + ".mp3"
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
