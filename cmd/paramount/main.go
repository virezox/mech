package main

import (
   "flag"
   "github.com/89z/format/dash"
   "github.com/89z/mech/paramount"
)

type downloader struct {
   dash.AdaptationSet
   info bool
   paramount.Media
}

func main() {
   var down downloader
   // b
   var guid string
   flag.StringVar(&guid, "b", "", "GUID")
   // d
   var isDASH bool
   flag.BoolVar(&isDASH, "d", false, "DASH download")
   // f
   // paramountplus.com/shows/video/622678414
   var bandwidth int64
   flag.Int64Var(&bandwidth, "f", 1622000, "target bandwidth")
   // i
   flag.BoolVar(&down.info, "i", false, "information")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if verbose {
      paramount.LogLevel = 1
   }
   if guid != "" {
      down.Media = paramount.NewMedia(guid)
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
