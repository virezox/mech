package main

import (
   "flag"
   "github.com/89z/mech/vimeo"
)

func main() {
   var clip vimeo.Clip
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   flag.Int64Var(&clip.ID, "b", 0, "clip ID")
   // c
   flag.Int64Var(&clip.UnlistedHash, "c", 0, "unlisted hash")
   // f
   var name string
   flag.StringVar(&name, "f", "720p", "public name")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      vimeo.LogLevel = 1
   }
   if clip.ID >= 1 || address != "" {
      err := doClip(&clip, address, name, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
