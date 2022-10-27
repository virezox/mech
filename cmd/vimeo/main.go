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
   verbose bool
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.Int64Var(&f.height, "f", 720, "target height")
   flag.BoolVar(&f.info, "i", false, "info only")
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      vimeo.Client.Log_Level = 2
   }
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
