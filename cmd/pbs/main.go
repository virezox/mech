package main

import (
   "flag"
   "github.com/89z/mech/pbs"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // f
   // http://pbs.org/wnet/nature/about-portugal-wild-land-edge
   var video int64
   flag.Int64Var(&video, "f", 2588259, "video bandwidth")
   // g
   // http://pbs.org/wnet/nature/about-portugal-wild-land-edge
   var audio string
   flag.StringVar(&audio, "g", "English", "audio name")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      pbs.LogLevel = 1
   }
   if address != "" {
      err := doWidget(address, audio, video, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
