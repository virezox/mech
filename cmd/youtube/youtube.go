package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
   "sort"
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

type video struct {
   address string
   audio string
   construct bool
   embed bool
   height int
   id string
   info bool
}

func (v video) do() error {
   play, err := v.player()
   if err != nil {
      return err
   }
   fmt.Println(play.Status())
   if v.info {
      fmt.Println(play.Details())
      for _, form := range play.StreamingData.AdaptiveFormats {
         str, err := form.WithURL("").Format()
         if err != nil {
            return err
         }
         fmt.Println(str)
      }
   } else {
      if v.height >= 1 {
         err := v.doVideo(play)
         if err != nil {
            return err
         }
      }
      if v.audio != "" {
         err := v.doAudio(play)
         if err != nil {
            return err
         }
      }
   }
   return nil
}

func (v video) doAudio(play *youtube.Player) error {
   for _, form := range play.StreamingData.AdaptiveFormats {
      if form.AudioQuality == v.audio {
         ext, err := form.Ext()
         if err != nil {
            return err
         }
         file, err := os.Create(play.Base() + ext)
         if err != nil {
            return err
         }
         defer file.Close()
         return form.Write(file)
      }
   }
   return nil
}

func (v video) doVideo(play *youtube.Player) error {
   sort.Sort(youtube.Height{play.StreamingData, v.height})
   for _, form := range play.StreamingData.AdaptiveFormats {
      ext, err := form.Ext()
      if err != nil {
         return err
      }
      file, err := os.Create(play.Base() + ext)
      if err != nil {
         return err
      }
      defer file.Close()
      if err := form.Write(file); err != nil {
         return err
      }
   }
   return nil
}
