package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/soundcloud"
   "time"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // s
   var sleep time.Duration
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   flag.Parse()
   if address != "" {
      tracks, err := soundcloud.Resolve(address)
      if err != nil {
         panic(err)
      }
      for _, track := range tracks {
         if info {
            fmt.Printf("%+v\n", track)
         } else {
            err := download(track)
            if err != nil {
               panic(err)
            }
            time.Sleep(sleep)
         }
      }
   } else {
      flag.Usage()
   }
}
