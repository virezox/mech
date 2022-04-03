package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
   "sort"
)

func (v video) doVideo(play *youtube.Player) error {
   if len(play.StreamingData.AdaptiveFormats) == 0 {
      return nil
   }
   form := play.StreamingData.AdaptiveFormats[0]
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

type video struct {
   address string
   audio string
   embed bool
   height int
   id string
   info bool
   token bool
}

func (v video) player() (*youtube.Player, error) {
   client := youtube.Android
   if v.embed {
      client = youtube.Embed
   }
   if v.id == "" {
      var err error
      v.id, err = youtube.VideoID(v.address)
      if err != nil {
         return nil, err
      }
   }
   if v.token {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      exc, err := youtube.OpenExchange(cache, "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      return client.PlayerHeader(exc.Header(), v.id)
   }
   return client.Player(v.id)
}

func (v video) do() error {
   play, err := v.player()
   if err != nil {
      return err
   }
   sort.SliceStable(play.StreamingData.AdaptiveFormats, func(int, int) bool {
      return true
   })
   sort.Stable(youtube.Height{play.StreamingData.AdaptiveFormats, v.height})
   if v.info {
      play.StreamingData.AdaptiveFormats.MediaType()
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

func doRefresh() error {
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

func doAccess() error {
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
