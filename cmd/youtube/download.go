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
   p, err := youtube.NewPlayer(id, auth, client)
   if err != nil {
      return err
   }
   if len(p.StreamingData.AdaptiveFormats) == 0 {
      return fmt.Errorf("%+v", p.PlayabilityStatus)
   }
   if len(c.formats) == 0 {
      c.formats["247"] = true
      c.formats["251"] = true
   }
   for _, form := range p.StreamingData.AdaptiveFormats {
      if c.formats[fmt.Sprint(form.Itag)] {
         exts, err := mech.ExtensionsByType(form.MimeType)
         if exts == nil {
            return fmt.Errorf("exts %v, err %v", exts, err)
         }
         name := p.Author() + "-" + p.Title() + exts[0]
         file, err := os.Create(strings.Map(mech.Clean, name))
         if err != nil {
            return err
         }
         defer file.Close()
         if err := form.Write(file); err != nil {
            // Videos can support both AdaptiveFormats and DASH:
            // zgJT91LA9gA
            return fmt.Errorf("%q %v", p.StreamingData.DashManifestURL, err)
         }
      }
   }
   return nil
}
