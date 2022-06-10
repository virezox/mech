package main

import (
   "flag"
   "github.com/89z/mech/paramount"
   "os"
   "path/filepath"
)

type downloader struct {
   clientPath string
   info bool
   keyPath string
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var d downloader
   // b
   var mediaID string
   flag.StringVar(&mediaID, "b", "", "media ID")
   // f
   var videoRate int64
   flag.Int64Var(&videoRate, "f", 1611000, "video bandwidth")
   // g
   var audioRate int64
   flag.Int64Var(&audioRate, "g", 999999, "audio bandwidth")
   // i
   flag.BoolVar(&d.info, "i", false, "information")
   // c
   d.clientPath = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&d.clientPath, "c", d.clientPath, "client ID")
   // k
   d.keyPath = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&d.keyPath, "k", d.keyPath, "private key")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // d
   var isDASH bool
   flag.BoolVar(&isDASH, "d", false, "DASH download")
   flag.Parse()
   if verbose {
      setVerbose()
   }
   ////////////////////////////////////////////////////////
   if mediaID != "" {
      var err error
      d.Preview, err = paramount.NewMedia(mediaID).Preview()
      if err != nil {
         panic(err)
      }
      if isDASH {
         err := d.DASH(videoRate, audioRate)
         if err != nil {
            panic(err)
         }
      } else {
         err := d.HLS(videoRate)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
