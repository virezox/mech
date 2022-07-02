package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/std/dash"
   "github.com/89z/std/mp4"
   "github.com/89z/std/os"
   "github.com/89z/mech/paramount"
   "github.com/89z/mech/widevine"
)

func (d downloader) DASH(video, audio int64) error {
   addr, err := paramount.New_Media(d.GUID).DASH()
   if err != nil {
      return err
   }
   res, err := paramount.Client.Get(addr.String())
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.url = addr
   if err := xml.NewDecoder(res.Body).Decode(&d.media); err != nil {
      return err
   }
   reps := d.media.Representations().Codecs("mp4a")
   if err := d.download(audio, reps); err != nil {
      return err
   }
   reps = d.media.Representations().Codecs("avc1")
   return d.download(video, reps)
}
