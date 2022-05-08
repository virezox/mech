package main

import (
   "flag"
   "github.com/89z/mech/bandcamp"
   "time"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // s
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bandcamp.LogLevel = 1
   }
   if address != "" {
      data, err := bandcamp.NewData(address)
      if err != nil {
         panic(err)
      }
      switch item.Item_Type {
      case "t", "a":
         err := arg.tralbum(item)
         if err != nil {
            panic(err)
         }
      case "i":
         err := arg.band(item)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
