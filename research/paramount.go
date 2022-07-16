package main

import (
   "flag"
   "github.com/89z/mech/research/mech"
   "github.com/89z/mech/paramount"
   "github.com/89z/rosso/dash"
   "os"
   "path/filepath"
   "strings"
)

type flags struct {
   bandwidth int
   codecs string
   dash bool
   guid string
   lang string
   mech.Flag
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.StringVar(&f.guid, "b", "", "GUID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // d
   flag.BoolVar(&f.dash, "d", false, "DASH download")
   // f
   flag.IntVar(&f.bandwidth, "f", 1611000, "video bandwidth")
   // g
   flag.StringVar(&f.codecs, "g", "mp4a", "audio codec")
   // h
   flag.StringVar(&f.lang, "h", "en", "audio lang")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   flag.Parse()
   if f.guid != "" {
      preview, err := paramount.New_Preview(f.guid)
      if err != nil {
         panic(err)
      }
      if err := f.DASH(preview); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) DASH(preview *paramount.Preview) error {
   address := paramount.DASH(f.guid)
   var err error
   f.Poster, err = paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   reps, err := f.Flag.DASH(address, preview.Base())
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      if r.MimeType != "audio/mp4" {
         return false
      }
      if r.Role() != "" {
         return false
      }
      return true
   })
   index := audio.Index(func(a, b dash.Representation) bool {
      if !strings.Contains(b.Codecs, f.codecs) {
         return false
      }
      if b.Adaptation.Lang != f.lang {
         return false
      }
      return true
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   return f.DASH_Get(video, video.Bandwidth(f.bandwidth))
}
