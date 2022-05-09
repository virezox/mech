package main

import (
   "flag"
   "github.com/89z/mech/roku"
)

func main() {
   // b
   var id string
   flag.StringVar(&id, "b", "", "ID")
   // dash
   var dash bool
   flag.BoolVar(&dash, "dash", false, "DASH")
   // f
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // k
   var key bool
   flag.BoolVar(&key, "key", false, "mp4decrypt key")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      roku.LogLevel = 1
   }
   if key {
      err := doKey(id)
      if err != nil {
         panic(err)
      }
   } else if id != "" {
      content, err := roku.NewContent(id)
      if err != nil {
         panic(err)
      }
      if dash {
         err := doDASH(content, bandwidth, info)
         if err != nil {
            panic(err)
         }
      } else {
         err := doHLS(content, bandwidth, info)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
