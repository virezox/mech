package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/ted"
)

func main() {
   // b
   var bitrate int64
   flag.Int64Var(&bitrate, "b", 180, "bitrate")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // s
   var slug string
   flag.StringVar(&slug, "s", "", "slug")
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
      fmt.Println("ted [flags]")
      flag.PrintDefaults()
   }
}
