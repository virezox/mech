package main

import (
   "flag"
   "github.com/89z/mech/paramount"
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
   var guid string
   flag.StringVar(&guid, "b", "", "GUID")
   // c
   down.client = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // d
   var isDASH bool
   flag.BoolVar(&isDASH, "d", false, "DASH download")
   // f
   // paramountplus.com/shows/video/x6XrF8A_tiSDRwc4Rt349KFKnCZ8QmtY
   var video int64
   flag.Int64Var(&video, "f", 1611000, "video bandwidth")
   // g
   // paramountplus.com/shows/video/x6XrF8A_tiSDRwc4Rt349KFKnCZ8QmtY
   var audio int64
   flag.Int64Var(&audio, "g", 999999, "audio bandwidth")
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
      paramount.LogLevel = 1
   }
   if guid != "" {
      var err error
      down.Preview, err = paramount.NewMedia(guid).Preview()
      if err != nil {
         panic(err)
      }
      if isDASH {
         err := down.DASH(video, audio)
         if err != nil {
            panic(err)
         }
      } else {
         err := down.HLS(video)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
