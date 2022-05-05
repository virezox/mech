package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/format/hls"
   "github.com/89z/mech/roku"
   "io"
   "net/http"
   "os"
)

func main() {
   // b
   var id string
   flag.StringVar(&id, "b", "", "ID")
   // f
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      roku.LogLevel = 1
   }
   if id != "" {
      err := doManifest(guid, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
