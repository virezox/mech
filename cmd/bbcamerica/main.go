package main

import (
   "flag"
   "github.com/89z/mech/bbcamerica"
   "os"
   "path/filepath"
)

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var down downloader
   // b
   var nid int64
   flag.Int64Var(&nid, "b", 0, "NID")
   // c
   down.client = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // f
   var video int64
   flag.Int64Var(&video, "f", 1920832, "video bandwidth")
   // i
   flag.BoolVar(&down.info, "i", false, "information")
   // k
   down.pem = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&down.pem, "k", down.pem, "private key")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      bbcamerica.LogLevel = 1
   }
   if nid >= 1 {
      err := down.doDASH(nid, video)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
