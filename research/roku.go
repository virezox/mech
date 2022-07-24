package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/roku"
   "github.com/89z/rosso/dash"
   "os"
   "path/filepath"
   "strings"
)

type flags struct {
   bandwidth int64
   codec string
   dash bool
   id string
   mech.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.StringVar(&f.id, "b", "", "ID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // d
   flag.BoolVar(&f.dash, "d", false, "DASH download")
   // f
   flag.Int64Var(&f.bandwidth, "f", 1920832, "video bandwidth")
   // g
   flag.StringVar(&f.codec, "g", "mp4a", "audio codec")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   flag.Parse()
   if f.id != "" {
      content, err := roku.New_Content(f.id)
      if err != nil {
         panic(err)
      }
      if err := f.DASH(content); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) DASH(content *roku.Content) error {
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   f.Poster, err = site.Playback(f.id)
   if err != nil {
      return err
   }
   f.Name = content.Name()
   reps, err := f.Stream.DASH(content.DASH().URL)
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
