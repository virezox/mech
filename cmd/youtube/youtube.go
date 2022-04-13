package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
   "sort"
)

func doRefresh() error {
   oauth, err := youtube.NewOAuth()
   if err != nil {
      return err
   }
   fmt.Println(oauth)
   fmt.Scanln()
   change, err := oauth.Exchange()
   if err != nil {
      return err
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   return change.Create(cache, "mech/youtube.json")
}

func doAccess() error {
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   change, err := youtube.OpenExchange(cache, "mech/youtube.json")
   if err != nil {
      return err
   }
   if err := change.Refresh(); err != nil {
      return err
   }
   return change.Create(cache, "mech/youtube.json")
}

func (v video) player() (*youtube.Player, error) {
   if v.id == "" {
      var err error
      v.id, err = youtube.VideoID(v.address)
      if err != nil {
         return nil, err
      }
   }
   if v.two {
      return youtube.AndroidEmbed.Player(v.id)
   }
   if v.three || v.four {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      change, err := youtube.OpenExchange(cache, "mech/youtube.json")
      if err != nil {
         return nil, err
      }
      if v.three {
         return youtube.AndroidRacy.Exchange(v.id, change)
      }
      return youtube.AndroidContent.Exchange(v.id, change)
   }
   return youtube.Android.Player(v.id)
}

func (v video) do() error {
   play, err := v.player()
   if err != nil {
      return err
   }
   forms := play.StreamingData.AdaptiveFormats
   if v.height >= 1 {
      sort.SliceStable(forms, func(int, int) bool {
         return true
      })
      sort.Stable(youtube.Height{forms, v.height})
   }
   if v.info {
      forms.MediaType()
      fmt.Println(play)
   } else {
      fmt.Println(play.Status())
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

func (v video) doVideo(play *youtube.Player) error {
   if len(play.StreamingData.AdaptiveFormats) == 0 {
      return nil
   }
   form := play.StreamingData.AdaptiveFormats[0]
   ext, err := mech.ExtensionByType(form.MimeType)
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

func (v video) doAudio(play *youtube.Player) error {
   for _, form := range play.StreamingData.AdaptiveFormats {
      if form.AudioQuality == v.audio {
         ext, err := mech.ExtensionByType(form.MimeType)
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
