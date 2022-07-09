package main

import (
   "github.com/89z/mech/amc"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Open_Auth(home + "/mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Create(home + "/mech/amc.json"); err != nil {
      return err
   }
   play, err := auth.Playback(f.nid)
   if err != nil {
      return err
   }
   rep, err := f.Representation(play.Source().Src, play.Base())
   if err != nil {
      return err
   }
   f.Poster = play
   if err := f.DASH(rep.Audio(), 0); err != nil {
      return err
   }
   video := rep.Video()
   return f.DASH(video, video.Bandwidth(f.bandwidth))
}
