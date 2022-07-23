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
   data := play.Data()
   f.Name = data.Get_Name()
   reps, err := f.DASH(data.Source().Src)
   if err != nil {
      return err
   }
   f.Poster = play
   if err := f.DASH_Get(reps.Audio(), 0); err != nil {
      return err
   }
   video := reps.Video()
   return f.DASH_Get(video, video.Bandwidth(f.bandwidth))
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
