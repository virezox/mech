package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/youtube"
   "os"
   "strings"
)

// Videos can support both AdaptiveFormats and DASH: zgJT91LA9gA
func (c choice) adaptiveFormats(play *youtube.Player) error {
   if len(c.itags) == 0 {
      c.itags = map[string]bool{
         "247": true, // youtube.com/watch?v=Leq8J0E2TQ0
         "251": true,
         "302": true, // youtube.com/watch?v=kVNl1P9StSU
      }
   }
   for i, form := range play.StreamingData.AdaptiveFormats {
      switch {
      case c.info:
         if i == 0 {
            fmt.Println(play.PlayabilityStatus)
            fmt.Println(play.VideoDetails)
         }
         fmt.Println(form)
      case c.itags[fmt.Sprint(form.Itag)]:
         name, err := filename(play, form)
         if err != nil {
            return err
         }
         file, err := os.Create(name)
         if err != nil {
            return err
         }
         defer file.Close()
         if err := form.Write(file); err != nil {
            return err
         }
      }
   }
   return nil
}

type choice struct {
   itags map[string]bool
   info bool
}

func filename(play *youtube.Player, form youtube.Format) (string, error) {
   name, err := format.ExtensionByType(form.MimeType)
   if err != nil {
      return "", err
   }
   name = play.VideoDetails.Author + "-" + play.VideoDetails.Title + name
   return strings.Map(format.Clean, name), nil
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

