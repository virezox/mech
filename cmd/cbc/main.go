package main

import (
   "flag"
)

type flags struct {
   audio_name string
   email string
   id string
   info bool
   password string
   video_bandwidth int
}

func main() {
   var f flags
   // b
   flag.StringVar(&f.id, "b", "", "ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // f
   flag.IntVar(&f.video_bandwidth, "f", 2052370, "video bandwidth")
   // g
   flag.StringVar(&f.audio_name, "g", "English", "audio name")
   // i
   flag.BoolVar(&f.info, "i", false, "information")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   if f.email != "" {
      err := f.profile()
      if err != nil {
         panic(err)
      }
   } else if f.id != "" {
      err := f.master()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
