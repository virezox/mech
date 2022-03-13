package main

import (
   "flag"
   "github.com/89z/mech/vimeo"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // f
   var name string
   flag.StringVar(&name, "f", "720p", "public name")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      vimeo.LogLevel = 1
   }
   if address != "" {
      err := doClip(address, name, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
