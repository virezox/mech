package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/github"
   "os"
)

func main() {
   var construct, exchange bool
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.Parse()
   if len(os.Args) == 1 {
      fmt.Println("github [flags]")
      flag.PrintDefaults()
      return
   }
   github.Verbose = true
   if exchange {
      err := authExchange()
      if err != nil {
         panic(err)
      }
   }
}
