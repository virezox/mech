package main

import (
   "flag"
   "github.com/89z/mech/abc"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // f
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 2_521_476, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      abc.LogLevel = 1
   }
   if address != "" {
      err := doManifest(address, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
