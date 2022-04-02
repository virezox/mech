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
   flag.BoolVar(&info, "i", false, "info only")
   // s
   var sleep time.Duration
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bandcamp.LogLevel = 1
   }
   if address != "" {
      item, err := bandcamp.NewItem(address)
      if err != nil {
         panic(err)
      }
      switch item.Item_Type {
      case "t", "a":
         err := doTralbum(item, info, sleep)
         if err != nil {
            panic(err)
         }
      case "i":
         err := doBand(item, info, sleep)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
