package main

import (
   "flag"
   "strings"
)

type flags struct {
   address string
   height int64
   info bool
}

func main() {
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "address")
   // f
   flag.Int64Var(&f.height, "f", 720, "target height")
   // i
   flag.BoolVar(&f.info, "i", false, "info only")
   flag.Parse()
   if strings.Contains(f.address, "vimeo.com/") {
      err := f.vimeo()
      if err != nil {
         panic(err)
      }
   } else if strings.Contains(f.address, "vhx.tv/") {
      err := f.vhx()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
