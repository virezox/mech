package main

import (
   "flag"
   "github.com/89z/format/dash"
   "github.com/89z/mech/paramount"
   "net/url"
   "os"
   "path/filepath"
)

type downloader struct {
   *dash.Period
   *paramount.Preview
   *url.URL
   client string
   info bool
   key []byte
   pem string
}

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
   // paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD
   var video int64
   flag.Int64Var(&video, "f", 2098819, "video bandwidth")
   // g
   // paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD
   var audio int64
   flag.Int64Var(&audio, "g", 131282, "audio bandwidth")
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
