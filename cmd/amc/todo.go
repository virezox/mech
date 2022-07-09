package main

import (
   "flag"
   "github.com/89z/mech/amc"
   "os"
   "path/filepath"
)

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

func (f flags) do() error {
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
   var p mech.Presentation
   p.Address = play.Source().Src
   p.Client_ID = f.client_ID
   p.Info = f.info
   p.Name = play.Base()
   p.Poster = play
   p.Private_Key
}
