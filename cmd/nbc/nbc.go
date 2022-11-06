package main

import (
   "github.com/89z/mech"
   "github.com/89z/mech/nbc"
)

type flags struct {
   bandwidth int64
   guid int64
   mech.Stream
   verbose bool
}

func (f flags) download() error {
   page, err := nbc.New_Bonanza_Page(f.guid)
   if err != nil {
      return err
   }
   video, err := page.Video()
   if err != nil {
      return err
   }
   f.Name = page.Name()
   master, err := f.HLS(video.ManifestPath)
   if err != nil {
      return err
   }
   streams := master.Streams
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}
