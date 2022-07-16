package main

import (
   "flag"
   //"github.com/89z/mech/research/pass"
   "github.com/89z/mech/research/fail"
   "github.com/89z/mech/paramount"
   "github.com/89z/rosso/dash"
   "os"
   "path/filepath"
)

type flags struct {
   bandwidth int
   codecs string
   dash bool
   guid string
   lang string
   //pass.Flag
   fail.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   flag.StringVar(&f.guid, "b", "", "GUID")
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   flag.BoolVar(&f.dash, "d", false, "DASH download")
   flag.IntVar(&f.bandwidth, "f", 1611000, "video bandwidth")
   flag.StringVar(&f.codecs, "g", "mp4a", "audio codec")
   flag.StringVar(&f.lang, "h", "en", "audio lang")
   flag.BoolVar(&f.Info, "i", false, "information")
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
   //reps, err := f.Flag.DASH(address, preview.Base())
   reps, err := f.Stream.DASH(address)
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      return r.MimeType == "audio/mp4"
   })
   return f.DASH_Get(audio, 0)
}
