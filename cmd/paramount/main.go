package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "os"
   "path/filepath"
)

type flags struct {
   bandwidth int
   codecs string
   dash bool
   guid string
   lang string
   mech.Stream
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
      preview, err := paramount.New_Media(f.guid).Preview()
      if err != nil {
         panic(err)
      }
      if f.dash {
         err := f.DASH(preview)
         if err != nil {
            panic(err)
         }
      } else {
         err := f.HLS(preview)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
