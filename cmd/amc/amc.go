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
   data := play.Data()
   self.Base = data.Base()
   reps, err := self.DASH(data.Source().Src)
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

func (self flags) login() error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(self.email, self.password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return auth.Create(home + "/mech/amc.json")
}
