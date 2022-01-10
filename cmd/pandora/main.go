package main

import (
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

func main() {
   var (
      info bool
      username, password string
   )
   flag.BoolVar(&info, "i", false, "information")
   flag.StringVar(&password, "p", "", "password")
   flag.StringVar(&username, "u", "", "username")
   flag.Parse()
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   cache = filepath.Join(cache, "/mech/pandora.json")
   switch {
   case username != "":
      err := login(cache, username, password)
      if err != nil {
         panic(err)
      }
   case flag.NArg() == 1:
      err := playback(cache, flag.Arg(0), info)
      if err != nil {
         panic(err)
      }
   default:
      fmt.Println("pandora [flags] [URL]")
      flag.PrintDefaults()
   }
}
