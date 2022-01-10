package main

import (
   "fmt"
   "github.com/89z/mech/pandora"
)

func playback(cache, addr string, info bool) error {
   user, err := pandora.OpenUserLogin(cache)
   if err != nil {
      return err
   }
   id, err := pandora.ID(addr)
   if err != nil {
      return err
   }
   play, err := user.PlaybackInfo(id)
   if err != nil {
      return err
   }
   if play.Result != nil {
      if info {
         fmt.Printf("%+v\n", play.Result.AudioUrlMap)
      } else {
         fmt.Println("download")
      }
   } else {
      fmt.Printf("%+v\n", play)
   }
   return nil
}

func login(username, password, cache string) error {
   part, err := pandora.NewPartnerLogin()
   if err != nil {
      return err
   }
   user, err := part.UserLogin(username, password)
   if err != nil {
      return err
   }
   fmt.Println("Create", cache)
   return user.Create(cache)
}
