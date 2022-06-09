package main

import (
   "flag"
   "github.com/89z/mech/apple"
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
   var contentID string
   flag.StringVar(&contentID, "b", "", "content ID")
   // c
   down.client = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // e
   var email string
   flag.StringVar(&email, "e", "", "email")
   // f
   var video int64
   flag.Int64Var(&video, "f", 1_999_999, "video bandwidth")
   // g
   var audio int64
   flag.Int64Var(&audio, "g", 127_000, "audio bandwidth")
   // i
   flag.BoolVar(&down.info, "i", false, "information")
   // k
   down.pem = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&down.pem, "k", down.pem, "private key")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      apple.LogLevel = 1
   }
   if email != "" {
      err := doLogin(email, password)
      if err != nil {
         panic(err)
      }
   } else if contentID != "" {
      err := down.doDASH(contentID, video, audio)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
