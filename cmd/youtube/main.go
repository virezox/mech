package main

import (
   "flag"
   "github.com/89z/mech/youtube"
   "strings"
)

func main() {
   var vid video
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
   // r
   var buf strings.Builder
   buf.WriteString("0: Android\n")
   buf.WriteString("1: Android embed\n")
   buf.WriteString("2: Android racy\n")
   buf.WriteString("3: Android content")
   flag.IntVar(&vid.request, "r", 0, buf.String())
   // refresh
   var refresh bool
   flag.BoolVar(&refresh, "refresh", false, "create OAuth refresh token")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      youtube.HTTP_Client.Level = 2
   }
   if refresh {
      err := do_refresh()
      if err != nil {
         panic(err)
      }
   } else if access {
      err := do_access()
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
