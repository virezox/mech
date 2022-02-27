package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

type choice struct {
   itags map[string]bool
   info bool
}

// Videos can support both AdaptiveFormats and DASH: zgJT91LA9gA
func (c choice) adaptiveFormats(play *youtube.Player) error {
   if play.PlayabilityStatus.Status != "OK" {
      fmt.Println(play.PlayabilityStatus)
   }
   if c.info {
      fmt.Println(play.VideoDetails)
   }
   if len(c.itags) == 0 {
      c.itags = map[string]bool{
         "247": true, // youtube.com/watch?v=Leq8J0E2TQ0
         "251": true,
         "302": true, // youtube.com/watch?v=kVNl1P9StSU
      }
   }
   for _, ada := range play.StreamingData.AdaptiveFormats {
      switch {
      case c.info:
         form, err := ada.Format(false)
         if err != nil {
            return err
         }
         fmt.Println(form)
      case c.itags[fmt.Sprint(ada.Itag)]:
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

func player(construct, embed bool, id string) (*youtube.Player, error) {
   client := youtube.Android
   if embed {
      client = youtube.Embed
   }
   if construct {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      ex, err := youtube.OpenExchange(cache + "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      return client.PlayerHeader(ex.Header(), id)
   }
   return client.Player(id)
}
