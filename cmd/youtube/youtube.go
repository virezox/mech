package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
)

func (v video) do() error {
   play, err := v.player()
   if err != nil {
      return err
   }
   if v.info {
      play.StreamingData.AdaptiveFormats.MediaType()
      fmt.Println(play)
   } else {
      fmt.Println(play.PlayabilityStatus)
      if v.height >= 1 {
         form := play.StreamingData.AdaptiveFormats.Video(v.height)
         err := download(form, play.Base())
         if err != nil {
            return err
         }
      }
      if v.audio != "" {
         form := play.StreamingData.AdaptiveFormats.Audio(v.audio)
         err := download(form, play.Base())
         if err != nil {
            return err
         }
      }
   }
   return nil
}

func download(form *youtube.Format, base string) error {
   ext, err := mech.ExtensionByType(form.MimeType)
   if err != nil {
      return err
   }
   file, err := os.Create(base + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := form.WriteTo(file); err != nil {
      return err
   }
   return nil
}

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

type video struct {
   address string
   audio string
   height int
   id string
   info bool
   two bool
   three bool
   four bool
}

