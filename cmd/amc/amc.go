package main

import (
   "github.com/89z/mech/amc"
   "os"
)

func (self flags) download() error {
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
   play, err := auth.Playback(self.nid)
   if err != nil {
      return err
   }
   self.Base = play.Base()
   reps, err := self.DASH(play.Source().Src)
   if err != nil {
      return err
   }
   self.Poster = play
   if err := self.DASH_Get(reps.Audio(), 0); err != nil {
      return err
   }
   video := reps.Video()
   return self.DASH_Get(video, video.Bandwidth(self.bandwidth))
}

func (f flags) login() error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(f.email, f.password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return auth.Create(home + "/mech/amc.json")
}
