package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/bandcamp"
   "time"
)

func main() {
   var choice flags
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   flag.BoolVar(&choice.info, "i", false, "info only")
   // s
   flag.DurationVar(&choice.sleep, "s", time.Second, "sleep")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bandcamp.LogLevel = 1
   }
   if address != "" {
      data, err := bandcamp.NewDataTralbum(address)
      if err != nil {
         panic(err)
      }
      if err := choice.process(data); err != nil {
         panic(err)
      }
   } else {
      fmt.Println("bandcamp [flags]")
      flag.PrintDefaults()
   }
}
