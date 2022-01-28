package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/bbc"
)

func main() {
   var (
      form int64
      info, verbose bool
   )
   flag.Int64Var(&form, "f", 0, "format")
   flag.BoolVar(&info, "i", false, "info")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bbc.LogLevel = 1
   }
   if flag.NArg() == 1 {
      addr := flag.Arg(0)
      item, err := bbc.NewNewsItem(addr)
      if err != nil {
         panic(err)
      }
      if err := media(item, info, form); err != nil {
         panic(err)
      }
   } else {
      fmt.Println("bbc [flags] [URL]")
      flag.PrintDefaults()
   }
}
