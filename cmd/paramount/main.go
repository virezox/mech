package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "os"
   "path/filepath"
)

type flags struct {
   dash bool
   guid string
   mech.Flags
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
   flag.IntVar(&f.Bandwidth_Video, "f", 1611000, "video bandwidth")
   // g
   flag.IntVar(&f.Bandwidth_Audio, "g", 999999, "audio bandwidth")
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
         err := f.do_DASH(preview)
         if err != nil {
            panic(err)
         }
      } else {
         err := f.do_HLS(preview)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
