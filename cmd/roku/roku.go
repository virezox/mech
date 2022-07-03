package main

import (
   "github.com/89z/mech/roku"
)

func (f flags) do_DASH(content *roku.Content) error {
   f.Address = content.DASH().URL
   site, err := roku.New_Cross_Site()
   if err != nil {
      return err
   }
   play, err := site.Playback(f.id)
   if err != nil {
      return err
   }
   return f.DASH(content.Base(), play)
}

func (f flags) do_HLS(content *roku.Content) error {
   video, err := content.HLS()
   if err != nil {
      return err
   }
   f.Address = video.URL
   return f.HLS(content.Base())
}
