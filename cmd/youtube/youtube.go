package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
   "strings"
)

func authConstruct(exc *youtube.Exchange) error {
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   fil, err := os.Open(cac + "/mech/youtube.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   return json.NewDecoder(fil).Decode(exc)
}

func authExchange() error {
   oau, err := youtube.NewOAuth()
   if err != nil {
      return err
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Press Enter to continue`, oau.Verification_URL, oau.User_Code)
   fmt.Scanln()
   exc, err := oau.Exchange()
   if err != nil {
      return err
   }
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cac += "/mech"
   os.Mkdir(cac, os.ModeDir)
   fil, err := os.Create(cac + "/youtube.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   enc := json.NewEncoder(fil)
   enc.SetIndent("", " ")
   return enc.Encode(exc)
}

func authRefresh() error {
   var exc youtube.Exchange
   err := authConstruct(&exc)
   if err != nil {
      return err
   }
   if err := exc.Refresh(); err != nil {
      return err
   }
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   fil, err := os.Create(cac + "/mech/youtube.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   enc := json.NewEncoder(fil)
   enc.SetIndent("", " ")
   return enc.Encode(exc)
}

func infoPath(play *youtube.Player, id string) error {
   fmt.Println("author:", play.Author())
   fmt.Println("title:", play.Title())
   fmt.Println()
   for _, form := range play.StreamingData.AdaptiveFormats {
      fmt.Println(form)
   }
   return nil
}

type choice map[string]bool

func (c choice) download(play *youtube.Player, id string) error {
   if len(c) == 0 {
      c["247"] = true // youtube.com/watch?v=Leq8J0E2TQ0
      c["251"] = true
      c["302"] = true // youtube.com/watch?v=kVNl1P9StSU
   }
   // Videos can support both AdaptiveFormats and DASH:
   // zgJT91LA9gA
   for _, form := range play.StreamingData.AdaptiveFormats {
      if c[fmt.Sprint(form.Itag)] {
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
