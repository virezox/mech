package main

import (
   "flag"
   "github.com/89z/mech/tiktok"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var awemeID int64
   flag.Int64Var(&awemeID, "b", 0, "aweme ID")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      tiktok.LogLevel = 1
   }
   if awemeID >= 1 || address != "" {
      err := detail(awemeID, address, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
