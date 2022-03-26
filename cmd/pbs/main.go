package main

import (
   "flag"
   "github.com/89z/mech/pbs"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      pbs.LogLevel = 1
   }
   if address != "" {
      err := doAsset(address, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
