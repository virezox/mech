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
   f.Name = page.Analytics.ConvivaAssetName
   master, err := f.HLS(video.ManifestPath)
   if err != nil {
      return err
   }
   streams := master.Streams
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}

type flags struct {
   bandwidth int64
   guid int64
   mech.Stream
}

func main() {
   var f flags
   // b
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   // f
   flag.Int64Var(&f.bandwidth, "f", 3_000_000, "target bandwidth")
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
