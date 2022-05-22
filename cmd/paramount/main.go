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
   *paramount.Preview
   dash.AdaptationSet
   info bool
   key []byte
   client string
   pem string
   *url.URL
}

func main() {
   cache, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var down downloader
   // b
   var guid string
   flag.StringVar(&guid, "b", "", "GUID")
   // c
   down.client = filepath.Join(cache, "mech/client_id.bin")
   flag.StringVar(&down.client, "c", down.client, "client ID")
   // d
   var isDASH bool
   flag.BoolVar(&isDASH, "d", false, "DASH download")
   // f
   // paramountplus.com/shows/video/622678414
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 1622000, "target bandwidth")
   // i
   flag.BoolVar(&down.info, "i", false, "information")
   // k
   down.pem = filepath.Join(cache, "mech/private_key.pem")
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
         err := down.DASH(bandwidth)
         if err != nil {
            panic(err)
         }
      } else {
         err := down.HLS(bandwidth)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
