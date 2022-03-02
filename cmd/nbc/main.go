package main

import (
   "flag"
   "github.com/89z/mech/nbc"
)

func main() {
   // f
   var form int64
   flag.Int64Var(&form, "f", 0, "format")
   // g
   var guid int64
   flag.Int64Var(&guid, "g", 0, "GUID")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      nbc.LogLevel = 1
   }
   if guid >= 1 {
      err := doManifest(guid, form, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
