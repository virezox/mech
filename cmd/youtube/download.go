package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
   "strings"
)

type choice struct {
   embed, construct bool
   formats map[string]bool
}

func (c choice) download(id string) error {
   auth := youtube.Key
   if c.construct {
      var exc youtube.Exchange
      err := authConstruct(&exc)
      if err != nil {
         return err
      }
      auth = youtube.Auth{"Authorization", "Bearer " + exc.Access_Token}
   }
   client := youtube.Android
   if c.embed {
      client = youtube.Embed
   }
   play, err := youtube.NewPlayer(id, auth, client)
   if err != nil {
      return err
   }
   if len(play.StreamingData.AdaptiveFormats) == 0 {
      return play.PlayabilityStatus
   }
   if len(c.formats) == 0 {
      c.formats["247"] = true
      c.formats["251"] = true
   }
   // Videos can support both AdaptiveFormats and DASH:
   // zgJT91LA9gA
   for _, form := range play.StreamingData.AdaptiveFormats {
      if c.formats[fmt.Sprint(form.Itag)] {
         ext, err := mech.ExtensionByType(form.MimeType)
         if err != nil {
            return err
         }
         name := play.Author() + "-" + play.Title() + ext
         file, err := os.Create(strings.Map(mech.Clean, name))
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
