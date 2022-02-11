package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/twitter"
)

func main() {
   // f
   var form int
   flag.IntVar(&form, "f", 0, "format")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info")
   // s
   var statusID int64
   flag.Int64Var(&statusID, "s", 0, "status ID")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      twitter.LogLevel = 1
   }
   if statusID >= 1 {
      err := statusPath(statusID, info, form)
      if err != nil {
         panic(err)
      }
   } else {
      fmt.Println("twitter [flags]")
      flag.PrintDefaults()
   }
}
