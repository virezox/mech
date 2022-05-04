package main

import (
   "flag"
   "github.com/89z/mech/cbc"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var id string
   flag.StringVar(&id, "b", "", "ID")
   // e
   var email string
   flag.StringVar(&email, "e", "", "email")
   // f
   // gem.cbc.ca/media/downton-abbey/s01e05
   var video int
   flag.IntVar(&video, "f", 2767506, "video bandwidth")
   // g
   // gem.cbc.ca/media/downton-abbey/s01e05
   var audio string
   flag.StringVar(&audio, "g", "English", "audio name")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "information")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      cbc.LogLevel = 1
   }
   if email != "" {
      err := doProfile(email, password)
      if err != nil {
         panic(err)
      }
   } else if id != "" || address != "" {
      err := doManifest(id, address, audio, video, info)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
