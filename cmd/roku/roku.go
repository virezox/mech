package main

import (
   "github.com/89z/mech/roku"
)

func (f flags) do_HLS(content *roku.Content) error {
   video, err := content.HLS()
   if err != nil {
      return err
   }
   master, err := f.HLS(video.URL, content.Base())
   if err != nil {
      return err
   }
   return f.HLS_Stream(master.Streams, 0)
}

func (f flags) do_DASH(content *roku.Content) error {
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   f.Poster, err = site.Playback(f.id)
   if err != nil {
      return err
   }
   reps, err := f.DASH(content.DASH().URL, content.Base())
   if err != nil {
      return err
   }
   return f.DASH_Get(reps, 0)
}
