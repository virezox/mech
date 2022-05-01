package main

import (
   "flag"
   "github.com/89z/mech/channel4"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var id string
   flag.StringVar(&id, "b", "", "program ID")
   // f
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 2_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // t
   var token string
   flag.StringVar(&token, "t", "", "token")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      channel4.LogLevel = 1
   }
   if id != "" || address != "" {
      err := doManifest(id, address, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else if token != "" {
      err := doToken(token)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
