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
   flag.BoolVar(&info, "i", false, "information")
   // s
   var sleep time.Duration
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   // v
   flag.Parse()
   if address != "" {
      tracks, err := soundcloud.Resolve(address)
      if err != nil {
         panic(err)
      }
      for i, track := range tracks {
         if info {
            fmt.Println(track)
         } else {
            if i >= 1 {
               time.Sleep(sleep)
            }
            err := download(track)
            if err != nil {
               panic(err)
            }
         }
      }
   } else {
      flag.Usage()
   }
}
