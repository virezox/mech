package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

func doExchange() error {
   oauth, err := youtube.NewOAuth()
   if err != nil {
      return err
   }
   fmt.Println(oauth)
   fmt.Scanln()
   exc, err := oauth.Exchange()
   if err != nil {
      return err
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   return exc.Create(cache, "/mech/youtube.json")
}

func doRefresh() error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   exc, err := youtube.OpenExchange(cache, "/mech/youtube.json")
   if err != nil {
      return err
   }
   if err := exc.Refresh(); err != nil {
      return err
   }
   return exc.Create(cache, "/mech/youtube.json")
}

func getID(videoID, address string) (string, error) {
   if videoID != "" {
      return videoID, nil
   }
   return youtube.VideoID(address)
}
