package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/research/dash"
   "strings"
)

func main() {
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "address")
   // f
   flag.Int64Var(&f.bandwidth, "f", 1, "video bandwidth")
   // g
   flag.StringVar(&f.codec, "g", "mp4a", "audio codec")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   flag.StringVar(&f.key, "k", "", "key")
   // o
   flag.StringVar(&f.Name, "o", "ballysports", "output")
   flag.Parse()
   if f.address != "" {
      if err := f.DASH(); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   address string
   bandwidth int64
   codec string
   key string
   mech.Stream
}

func (f flags) DASH() error {
   reps, err := f.Stream.DASH(f.address)
   if err != nil {
      return err
   }
   audio := reps.Audio()
   index := audio.Index(func(a, b dash.Representation) bool {
      return strings.Contains(b.Codecs, f.codec)
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   return f.DASH_Get(video, video.Bandwidth(f.bandwidth))
}
