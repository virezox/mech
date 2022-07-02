package main

import (
   "encoding/xml"
   "fmt"
   "github.com/89z/mech/amc"
   "github.com/89z/std/dash"
   "os"
)

func (f flags) DASH() error {
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
   if f.nid == 0 {
      f.nid, err = amc.Get_NID(f.address)
      if err != nil {
         return err
      }
   }
   play, err := auth.Playback(f.nid)
   if err != nil {
      return err
   }
   source := play.Source()
   res, err := amc.Client.Redirect(nil).Get(source.Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var media dash.Media
   if err := xml.NewDecoder(res.Body).Decode(&media); err != nil {
      return err
   }
   reps := media.Representations().Video()
   if f.video_bandwidth >= 1 {
      rep := reps.Get_Bandwidth(f.video_bandwidth)
      if f.info {
         for _, each := range reps {
            if each.Bandwidth == rep.Bandwidth {
               fmt.Print("!")
            }
            fmt.Println(each)
         }
      } else {
         var key []byte
         if source.Key_Systems != nil {
            key, err = f.key(play, rep.ContentProtection.Default_KID)
            if err != nil {
               return err
            }
         }
         return download(rep, key, play.Base())
      }
   }
   return nil
}
