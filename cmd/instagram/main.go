package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/instagram"
)

func main() {
   // a
   var auth bool
   flag.BoolVar(&auth, "a", false, "authentication")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // u
   var username string
   flag.StringVar(&username, "u", "", "username")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      instagram.LogLevel = 1
   }
   if username != "" {
      err := saveLogin(username, password)
      if err != nil {
         panic(err)
      }
   } else if flag.NArg() == 1 {
      shortcode := flag.Arg(0)
      if !instagram.Valid(shortcode) {
         panic("invalid shortcode")
      }
      if auth {
         err := doItems(shortcode, info)
         if err != nil {
            panic(err)
         }
      } else {
         err := doGraph(shortcode)
         if err != nil {
            panic(err)
         }
      }
   } else {
      fmt.Println("instagram [flags] [shortcode]")
      flag.PrintDefaults()
   }
}
