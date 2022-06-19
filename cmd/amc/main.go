package main

import (
   "flag"
   "github.com/89z/mech/amc"
   "github.com/89z/format/dash"
   "net/url"
   "os"
   "path/filepath"
)

type downloader struct {
   *amc.Playback
   client string
   info bool
   key []byte
   pem string
   url *url.URL
   media dash.Media
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var down downloader
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var nid int64
   flag.Int64Var(&nid, "b", 0, "NID")
   // c
   down.client = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // e
   var email string
   flag.StringVar(&email, "e", "", "email")
   // f
   // amcplus.com/shows/orphan-black/episodes/season-1-natural-selection--1011153
   var video int64
   flag.Int64Var(&video, "f", 1_999_999, "video bandwidth")
   // g
   // amcplus.com/shows/orphan-black/episodes/season-1-natural-selection--1011153
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
      amc.LogLevel = 1
   }
   if email != "" {
      err := do_login(email, password)
      if err != nil {
         panic(err)
      }
   } else if nid >= 1 || address != "" {
      err := down.do_DASH(address, nid, video, audio)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
