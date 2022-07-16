package main

import (
   "github.com/89z/mech/roku"
   "github.com/89z/rosso/dash"
   "strings"
)

func (self flags) DASH(content *roku.Content) error {
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   self.Poster, err = site.Playback(self.id)
   if err != nil {
      return err
   }
   self.Base = content.Base()
   reps, err := self.Stream.DASH(content.DASH().URL)
   if err != nil {
      return err
   }
   audio := reps.Audio()
   index := audio.Index(func(a, b dash.Representation) bool {
      return strings.Contains(b.Codecs, self.codec)
   })
   if err := self.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   return self.DASH_Get(video, video.Bandwidth(self.bandwidth))
}

func (self flags) HLS(content *roku.Content) error {
   video, err := content.HLS()
   if err != nil {
      return err
   }
   self.Base = content.Base()
   master, err := self.Stream.HLS(video.URL)
   if err != nil {
      return err
   }
   streams := master.Streams
   return self.HLS_Streams(streams, streams.Bandwidth(self.bandwidth))
}
