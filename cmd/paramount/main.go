package main

import (
   "flag"
   "github.com/89z/format"
   "github.com/89z/mech/paramount"
)

func main() {
   // b
   var guid string
   flag.StringVar(&guid, "b", "", "GUID")
   // f
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 3_063_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var logLevel int
   flag.IntVar(&logLevel, "v", 0, "log level")
   flag.Parse()
   paramount.LogLevel = format.LogLevel(logLevel)
   if guid != "" {
      err := doManifest(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
