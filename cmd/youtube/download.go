package main

import (
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
   "strings"
)

func downloadPath(construct, embed bool, atag, vtag int, id string) error {
   auth := youtube.Key
   if construct {
      var exc youtube.Exchange
      err := authConstruct(&exc)
      if err != nil {
         return err
      }
      auth = youtube.Auth{"Authorization", "Bearer " + exc.Access_Token}
   }
   client := youtube.Android
   if embed {
      client = youtube.Embed
   }
   p, err := youtube.NewPlayer(id, auth, client)
   if err != nil {
      return err
   }
   if p.StreamingData.DashManifestURL != "" {
      return mech.Invalid{p.StreamingData.DashManifestURL}
   }
   formats := []youtube.Format{
      {Itag: atag}, {Itag: vtag, Height: 720},
   }
   for _, a := range formats {
      var keep func(youtube.Format)bool
      switch a.Itag {
      case -1:
         continue
      case 0:
         keep = func(b youtube.Format) bool {
            return b.Height <= a.Height
         }
      default:
         keep = func(b youtube.Format) bool {
            return b.Itag == a.Itag
         }
      }
      var forms []youtube.Format
      for _, form := range p.StreamingData.AdaptiveFormats {
         if keep(form) {
            forms = append(forms, form)
         }
      }
      if forms != nil {
         return download(p, forms[0])
      }
      fmt.Printf("%+v\n", p.PlayabilityStatus)
   }
   return nil
}

func download(p *youtube.Player, f youtube.Format) error {
   exts, err := mech.ExtensionsByType(f.MimeType)
   if exts == nil {
      return fmt.Errorf("exts %v, err %v", exts, err)
   }
   name := p.Author() + "-" + p.Title() + exts[0]
   file, err := os.Create(strings.Map(mech.Clean, name))
   if err != nil {
      return err
   }
   defer file.Close()
   return f.Write(file)
}
