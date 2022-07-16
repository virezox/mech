package main

import (
   "github.com/89z/mech/paramount"
   "github.com/89z/rosso/dash"
   "github.com/89z/rosso/hls"
   "strings"
)

func (self flags) DASH(preview *paramount.Preview) error {
   var err error
   self.Poster, err = paramount.New_Session(self.guid)
   if err != nil {
      return err
   }
   self.Base = preview.Base()
   reps, err := self.Stream.DASH(paramount.DASH(self.guid))
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      if r.MimeType != "audio/mp4" {
         return false
      }
      if r.Role() == "decription" {
         return false
      }
      return true
   })
   index := audio.Index(func(a, b dash.Representation) bool {
      if !strings.HasPrefix(b.Adaptation.Lang, self.lang) {
         return false
      }
      if !strings.HasPrefix(b.Codecs, self.codecs) {
         return false
      }
      return true
   })
   if err := self.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   return self.DASH_Get(video, video.Bandwidth(self.bandwidth))
}

func (self flags) HLS(preview *paramount.Preview) error {
   self.Base = preview.Base()
   master, err := self.Stream.HLS(paramount.HLS(self.guid))
   if err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   return self.HLS_Streams(streams, streams.Bandwidth(self.bandwidth))
}
