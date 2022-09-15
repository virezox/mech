package main

import (
   "flag"
   "github.com/89z/mech/vimeo"
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
   } else if vimeo.Is_Embed(f.address) {
      err := f.vhx()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
