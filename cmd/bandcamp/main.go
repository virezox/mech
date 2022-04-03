package main

import (
   "flag"
   "github.com/89z/mech/bandcamp"
   "time"
)

func main() {
   var arg argument
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // e
   flag.StringVar(&arg.ext, "e", ".mp3", "extension")
   // i
   flag.BoolVar(&arg.info, "i", false, "information")
   // s
   flag.DurationVar(&arg.sleep, "s", time.Second, "sleep")
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
