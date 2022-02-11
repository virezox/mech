package main

import (
   "flag"
   "github.com/89z/mech/bbc"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // f
   var form int64
   flag.Int64Var(&form, "f", 0, "format")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bbc.LogLevel = 1
   }
   if address != "" {
      item, err := bbc.NewNewsItem(address)
      if err != nil {
         panic(err)
      }
      if err := media(item, info, form); err != nil {
         panic(err)
      }
   } else {
      flag.PrintDefaults()
   }
}
