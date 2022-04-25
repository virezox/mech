package main

import (
   "flag"
   "github.com/89z/mech/nbc"
)

func main() {
   // b
   var guid int64
   flag.Int64Var(&guid, "b", 0, "GUID")
   // f
   var bandwidth int
   // nbc.com/saturday-night-live/video/march-12-zoe-kravitz/9000199371
   // nbc.com/saturday-night-live/video/may-15-keeganmichael-key/4358937
   flag.IntVar(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      nbc.LogLevel = 1
   }
   if guid >= 1 {
      err := doManifest(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
