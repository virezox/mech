package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/pandora"
   "os"
   "path/filepath"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
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
      pandora.LogLevel = 1
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "/mech/pandora.json")
   if username != "" {
      err := login(cache, username, password)
      if err != nil {
         panic(err)
      }
   } else if address != "" {
      err := playback(cache, address, info)
      if err != nil {
         panic(err)
      }
   } else {
      fmt.Println("pandora [flags]")
      flag.PrintDefaults()
   }
}
