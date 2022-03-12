package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

func (v video) do() error {
   client := youtube.Android
   if v.embed {
      client = youtube.Embed
   }
   var (
      err error
      play *youtube.Player
   )
   if v.id == "" {
      v.id, err = youtube.VideoID(v.address)
      if err != nil {
         return err
      }
   }
   if v.construct {
      cache, err := os.UserCacheDir()
      if err != nil {
         return err
      }
      exc, err := youtube.OpenExchange(cache, "/mech/youtube.json")
      if err != nil {
         return err
      }
      play, err = client.PlayerHeader(exc.Header(), v.id)
   } else {
      play, err = client.Player(v.id)
   }
   if err != nil {
      return err
   }
   fmt.Println(play.Status())
   if v.info {
      fmt.Println(play.Details())
      for _, form := range play.StreamingData.AdaptiveFormats {
         form.URL = ""
         str, err := form.Format()
         if err != nil {
            return err
         }
         fmt.Println(str)
      }
   } else {
      if v.height >= 1 {
         name, err := ada.Name(play)
         if err != nil {
            return err
         }
         file, err := os.Create(name)
         if err != nil {
            return err
         }
         defer file.Close()
         if err := ada.Write(file); err != nil {
            return err
         }
      }
      if v.audio != "" {
         name, err := ada.Name(play)
         if err != nil {
            return err
         }
         file, err := os.Create(name)
         if err != nil {
            return err
         }
         defer file.Close()
         if err := ada.Write(file); err != nil {
            return err
         }
      }
   }
   return nil
}

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
