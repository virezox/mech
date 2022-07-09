package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/nbc"
)

func (f flags) download() error {
   page, err := nbc.New_Bonanza_Page(f.guid)
   if err != nil {
      return err
   }
   video, err := page.Video()
   if err != nil {
      return err
   }
   master, err := f.HLS(video.ManifestPath, page.Analytics.ConvivaAssetName)
   if err != nil {
      return err
   }
   return f.HLS_Stream(master.Streams, 0)
}

type flags struct {
   bandwidth int
   guid int64
   mech.Flag
}

func main() {
   var f flags
   // b
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   // f
   flag.IntVar(&f.bandwidth, "f", 3_000_000, "target bandwidth")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.Parse()
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
