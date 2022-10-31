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
   verbose bool
   video_ID string
}

func main() {
   var f flags
   // b
   flag.StringVar(&f.video_ID, "b", "", "video ID")
   // f
   flag.IntVar(&f.height, "f", 1080, "target video height")
   // g
   flag.StringVar(&f.audio, "g", "AUDIO_QUALITY_MEDIUM", "target audio")
   // i
   flag.BoolVar(&f.info, "i", false, "information")
   // refresh
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   // access
   flag.BoolVar(&f.access, "access", false, "create OAuth access token")
   // r
   var buf strings.Builder
   buf.WriteString("0: Android\n")
   buf.WriteString("1: Android embed\n")
   buf.WriteString("2: Android racy\n")
   buf.WriteString("3: Android content")
   flag.IntVar(&f.request, "r", 0, buf.String())
   // a
   flag.Func("a", "address", func(s string) error {
      return youtube.Video_ID(s, &f.video_ID)
   })
   // v
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      youtube.HTTP_Client.Log_Level = 2
   }
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
