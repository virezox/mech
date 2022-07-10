package main

import (
   "flag"
   "github.com/89z/mech/youtube"
   "strings"
)

type flags struct {
   access bool
   audio string
   height int
   info bool
   refresh bool
   request int
   video_ID string
}

func main() {
   var f flags
   // b
   flag.StringVar(&f.video_ID, "b", "", "video ID")
   // f
   flag.IntVar(&f.height, "f", 720, "target video height")
   // g
   flag.StringVar(&f.audio, "g", "AUDIO_QUALITY_MEDIUM", "target audio")
   // i
   flag.BoolVar(&f.info, "i", false, "information")
   // refresh
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   // access
   flag.BoolVar(&f.access, "access", false, "create OAuth access token")
   // r
   var b strings.Builder
   b.WriteString("0: Android\n")
   b.WriteString("1: Android embed\n")
   b.WriteString("2: Android racy\n")
   b.WriteString("3: Android content")
   flag.IntVar(&f.request, "r", 0, b.String())
   // a
   flag.Func("a", "address", func(s string) error {
      return youtube.Video_ID(s, &f.video_ID)
   })
   flag.Parse()
   if f.refresh {
      err := refresh()
      if err != nil {
         panic(err)
      }
   } else if f.access {
      err := access()
      if err != nil {
         panic(err)
      }
   } else if f.video_ID != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
