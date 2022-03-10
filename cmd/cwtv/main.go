package main

import (
   "flag"
   "github.com/89z/mech/cwtv"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // f
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 1_999_999, "min bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      cwtv.LogLevel = 1
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
