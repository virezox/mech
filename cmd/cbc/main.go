package main

import (
   "flag"
   "github.com/89z/mech/cbc"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var id string
   flag.StringVar(&id, "b", "", "ID")
   // e
   var email string
   flag.StringVar(&email, "e", "", "email")
   // f
   var bandwidth int
   flag.IntVar(&bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      cbc.LogLevel = 1
   }
   if email != "" {
      err := doProfile(email, password)
      if err != nil {
         panic(err)
      }
   } else if id != "" || address != "" {
      err := doManifest(id, address, bandwidth, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
