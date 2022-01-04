package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/twitter"
)

func main() {
   var (
      info, space, verbose bool
      format int
   )
   flag.BoolVar(&info, "i", false, "info")
   flag.IntVar(&format, "f", 0, "format")
   flag.BoolVar(&space, "s", false, "space")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      twitter.Log.Level = 1
   }
   if flag.NArg() == 1 {
      id := flag.Arg(0)
      if space {
         err := spacePath(id, info)
         if err != nil {
            panic(err)
         }
      } else {
         err := statusPath(id, info, format)
         if err != nil {
            panic(err)
         }
      }
   } else {
      fmt.Println("twitter [flags] [ID]")
      flag.PrintDefaults()
   }
}
