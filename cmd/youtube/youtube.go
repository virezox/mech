package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/youtube"
   "net/http"
   "os"
   "strings"
)

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
      return ex.Player(client, id)
   }
   res, err := youtube.NewPlayerResponse(client, id)
   if err != nil {
      return nil, err
   }
   return res.Player()
}

type choice struct {
   itags map[string]bool
   info bool
   useFormats bool
}

func filename(play *youtube.Player, form youtube.Format) (string, error) {
   name, err := format.ExtensionByType(form.MimeType)
   if err != nil {
      return "", err
   }
   name = play.VideoDetails.Author + "-" + play.VideoDetails.Title + name
   return strings.Map(format.Clean, name), nil
}

// Videos can support both AdaptiveFormats and DASH: zgJT91LA9gA
func (c choice) adaptiveFormats(play *youtube.Player) error {
   if len(c.itags) == 0 {
      c.itags = map[string]bool{
         "247": true, // youtube.com/watch?v=Leq8J0E2TQ0
         "251": true,
         "302": true, // youtube.com/watch?v=kVNl1P9StSU
      }
   }
   for _, form := range play.StreamingData.AdaptiveFormats {
      if c.info {
         fmt.Println(form)
      } else if c.itags[fmt.Sprint(form.Itag)] {
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

func (c choice) formats(play *youtube.Player) error {
   for _, form := range play.StreamingData.Formats {
      if c.info {
         fmt.Println(form)
      } else if c.itags[fmt.Sprint(form.Itag)] {
         os.Stdout.WriteString("GET ")
         format.Trim(os.Stdout, form.URL)
         os.Stdout.WriteString("\n")
         res, err := http.Get(form.URL)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         name, err := filename(play, form)
         if err != nil {
            return err
         }
         file, err := os.Create(name)
         if err != nil {
            return err
         }
         defer file.Close()
         pro := format.NewProgress(res, os.Stdout)
         if _, err := file.ReadFrom(pro); err != nil {
            return err
         }
      }
   }
   return nil
}
