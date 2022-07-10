package main

import (
   "flag"
   "github.com/89z/mech"
   "os"
   "path/filepath"
)

type flags struct {
   content_ID string
   mech.Flag
   bandwidth int
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.StringVar(&f.content_ID, "b", "", "content ID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // f
   flag.IntVar(&f.bandwidth, "f", 1920832, "video bandwidth")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   flag.Parse()
   if f.content_ID != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
