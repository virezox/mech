package main

import (
   "flag"
   "github.com/89z/mech/vimeo"
)

type flags struct {
   address string
   height int64
   info bool
   password string
   verbose bool
}

func main() {
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "address")
   // f
   flag.Int64Var(&f.height, "f", 720, "target height")
   // i
   flag.BoolVar(&f.info, "i", false, "info only")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   // v
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      vimeo.Client.Log_Level = 2
   }
   if f.address != "" {
      clip, err := vimeo.New_Clip(f.address)
      if err != nil {
         panic(err)
      }
      if f.password != "" {
         err := f.auth(clip)
         if err != nil {
            panic(err)
         }
      } else {
         err := f.anon(clip)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
