package main

import (
   "flag"
   "github.com/89z/mech/youtube"
)

type video struct {
   address string
   audio string
   height int
   id string
   info bool
   two bool
   three bool
   four bool
}

func main() {
   var vid video
   // 2
   flag.BoolVar(&vid.two, "2", false, "request type two")
   // 3
   flag.BoolVar(&vid.three, "3", false, "request type three")
   // 4
   flag.BoolVar(&vid.four, "4", false, "request type four")
   // a
   flag.StringVar(&vid.address, "a", "", "address")
   // access
   var access bool
   flag.BoolVar(&access, "access", false, "create OAuth access token")
   // b
   flag.StringVar(&vid.id, "b", "", "video ID")
   // f
   flag.IntVar(&vid.height, "f", 720, "target video height")
   // g
   flag.StringVar(&vid.audio, "g", "AUDIO_QUALITY_MEDIUM", "target audio")
   // i
   flag.BoolVar(&vid.info, "i", false, "information")
   // refresh
   var refresh bool
   flag.BoolVar(&refresh, "refresh", false, "create OAuth refresh token")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      youtube.LogLevel = 1
   }
   if refresh {
      err := doRefresh()
      if err != nil {
         panic(err)
      }
   } else if access {
      err := doAccess()
      if err != nil {
         panic(err)
      }
   } else if vid.id != "" || vid.address != "" {
      err := vid.do()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
