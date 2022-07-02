package main

import (
   "flag"
   "github.com/89z/mech"
   "os"
   "path/filepath"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.Int64Var(&f.nid, "b", 0, "NID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // f
   flag.IntVar(&f.Bandwidth_Video, "f", 1_999_999, "video bandwidth")
   // g
   flag.IntVar(&f.Bandwidth_Audio, "g", 127_000, "audio bandwidth")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   if f.email != "" {
      err := f.login()
      if err != nil {
         panic(err)
      }
   } else if f.nid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   mech.Flags
   email string
   nid int64
   password string
}
