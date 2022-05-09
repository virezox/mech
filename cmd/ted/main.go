package main

import (
   "flag"
   "github.com/89z/mech/ted"
)

func main() {
   // b
   var slug string
   flag.StringVar(&slug, "b", "", "slug")
   // f
   var bitrate int64
   flag.Int64Var(&bitrate, "f", 180, "bitrate")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      ted.LogLevel = 1
   }
   if slug != "" {
      err := process(slug, info, bitrate)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
