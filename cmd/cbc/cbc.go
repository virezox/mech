package main

import (
   "github.com/89z/mech/cbc"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   profile, err := cbc.Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      return err
   }
   asset, err := cbc.New_Asset(f.id)
   if err != nil {
      return err
   }
   asset_media, err := profile.Media(asset)
   if err != nil {
      return err
   }
   master, err := f.HLS(*asset_media.URL, asset.AppleContentId)
   if err != nil {
      return err
   }
   if err := f.HLS_Media(master.Media, 0); err != nil {
      return err
   }
   return f.HLS_Stream(master.Streams, 0)
}

func (f flags) profile() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   login, err := cbc.New_Login(f.email, f.password)
   if err != nil {
      return err
   }
   web, err := login.Web_Token()
   if err != nil {
      return err
   }
   top, err := web.Over_The_Top()
   if err != nil {
      return err
   }
   profile, err := top.Profile()
   if err != nil {
      return err
   }
   return profile.Create(home + "/mech/cbc.json")
}
