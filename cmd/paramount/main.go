package main

import (
   "flag"
   "github.com/89z/mech/paramount"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var guid string
   flag.StringVar(&guid, "b", "", "GUID")
   // f
   // paramountplus.com/shows/video/186A7772-3307-4270-4FB9-D16F4DC7E08C
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 2_223_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      paramount.LogLevel = 1
   }
   if guid != "" || address != "" {
      err := newMaster(guid, address, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
