package main

import (
   "flag"
   "github.com/89z/mech"
   "github.com/89z/mech/nbc"
)

func (self flags) download() error {
   page, err := nbc.New_Bonanza_Page(self.guid)
   if err != nil {
      return err
   }
   video, err := page.Video()
   if err != nil {
      return err
   }
   self.Base = page.Analytics.ConvivaAssetName
   master, err := self.HLS(video.ManifestPath)
   if err != nil {
      return err
   }
   streams := master.Streams
   return self.HLS_Streams(streams, streams.Bandwidth(self.bandwidth))
}

type flags struct {
   bandwidth int
   guid int64
   mech.Stream
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
