package main

import (
   "github.com/89z/mech/roku"
   "github.com/89z/rosso/dash"
   "strings"
)

func (f flags) DASH(content *roku.Content) error {
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   f.Poster, err = site.Playback(f.id)
   if err != nil {
      return err
   }
   f.Name = content.Name()
   reps, err := f.Stream.DASH(content.DASH().URL)
   if err != nil {
      return err
   }
   audio := reps.Audio()
   index := audio.Index(func(a, b dash.Representation) bool {
      return strings.Contains(b.Codecs, f.codec)
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   return f.DASH_Get(video, video.Bandwidth(f.bandwidth))
}

func (f flags) HLS(content *roku.Content) error {
   video, err := content.HLS()
   if err != nil {
      return err
   }
   f.Name = content.Name()
   master, err := f.Stream.HLS(video.URL)
   if err != nil {
      return err
   }
   streams := master.Streams
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}
