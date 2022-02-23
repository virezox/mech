package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/soundcloud"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   flag.Parse()
   if address != "" {
      track, err := soundcloud.Resolve(address)
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
      flag.PrintDefaults()
   }
}
