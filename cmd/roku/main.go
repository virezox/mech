package main

import (
   "flag"
   "github.com/89z/mech/roku"
   "github.com/89z/mech"
   "os"
   "path/filepath"
)

type flags struct {
   dash bool
   id string
   mech.Flags
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
   flag.IntVar(&f.Bandwidth_Video, "f", 1920832, "video bandwidth")
   // g
   flag.IntVar(&f.Bandwidth_Audio, "g", 128000, "audio bandwidth")
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
      if f.dash {
         err := f.do_DASH(content)
         if err != nil {
            panic(err)
         }
      } else {
         err := f.do_HLS(content)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
