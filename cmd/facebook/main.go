package main

import (
   "flag"
   "github.com/89z/mech/facebook"
)

func main() {
   // b
   var videoID int64
   flag.Int64Var(&videoID, "b", 0, "video ID")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      facebook.LogLevel = 1
   }
   if videoID >= 1 {
      err := doVideo(videoID, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
