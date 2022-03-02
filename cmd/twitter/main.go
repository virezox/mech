package main

import (
   "flag"
   "github.com/89z/mech/twitter"
)

func main() {
   // b
   var statusID int64
   flag.Int64Var(&statusID, "b", 0, "status ID")
   // bitrate
   var bitrate int64
   flag.Int64Var(&bitrate, "bitrate", 2_176_000, "bitrate")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      twitter.LogLevel = 1
   }
   if statusID >= 1 {
      err := statusPath(statusID, bitrate, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
