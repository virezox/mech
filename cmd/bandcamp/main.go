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
      param, err := bandcamp.NewParams(address)
      if err != nil {
         panic(err)
      }
      if param.I_Type != "" {
         tralb, err := param.Tralbum()
         if err != nil {
            panic(err)
         }
         if err := doTralbum(tralb, info, sleep); err != nil {
            panic(err)
         }
      } else {
         err := doBand(param, info, sleep)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
