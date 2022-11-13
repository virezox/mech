package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/paramount"
   "github.com/89z/mech/widevine"
   "os"
   "path/filepath"
)

type flags struct {
   codecs string
   dash bool
   guid string
   height int64
   lang string
   mech.Stream
   verbose bool
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
   flag.Int64Var(&f.height, "f", 720, "video height")
   // g
   flag.StringVar(&f.codecs, "g", "mp4a", "audio codec")
   // h
   flag.StringVar(&f.lang, "h", "en", "audio lang")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   // v
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      paramount.Client.Log_Level = 2
      widevine.Client.Log_Level = 2
   }
   if f.guid != "" {
      preview, err := paramount.New_Preview(f.guid)
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
