package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech"
   "github.com/89z/mech/youtube"
   "os"
)

func (v video) do() error {
   play, err := v.player()
   if err != nil {
      return err
   }
   forms := play.StreamingData.AdaptiveFormats
   if v.info {
      text, err := play.MarshalText()
      if err != nil {
         return err
      }
      os.Stdout.Write(text)
   } else {
      fmt.Println(play.PlayabilityStatus)
      if v.audio != "" {
         form, ok := forms.Audio(v.audio)
         if ok {
            err := download(form, play.Base())
            if err != nil {
               return err
            }
         }
      }
      if v.height >= 1 {
         form, ok := forms.Video(v.height)
         if ok {
            err := download(form, play.Base())
            if err != nil {
               return err
            }
         }
      }
   }
   return nil
}

func do_refresh() error {
   auth, err := youtube.New_OAuth()
   if err != nil {
      return err
   }
   fmt.Println(auth)
   fmt.Scanln()
   change, err := auth.Exchange()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return change.Create(home + "/mech/youtube.json")
}

func do_access() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   change, err := youtube.Open_Exchange(home + "/mech/youtube.json")
   if err != nil {
      return err
   }
   if err := change.Refresh(); err != nil {
      return err
   }
   return change.Create(home + "/mech/youtube.json")
}
func download(form *youtube.Format, base string) error {
   ext, err := mech.Extension_By_Type(form.MimeType)
   if err != nil {
      return err
   }
   file, err := format.Create(base + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   return form.Encode(file)
}

func (v video) player() (*youtube.Player, error) {
   if v.id == "" {
      var err error
      v.id, err = youtube.Video_ID(v.address)
      if err != nil {
         return nil, err
      }
   }
   if v.request == 1 {
      return youtube.Android_Embed.Player(v.id)
   }
   if v.request >= 2 {
      home, err := os.UserHomeDir()
      if err != nil {
         return nil, err
      }
      change, err := youtube.Open_Exchange(home + "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      if v.request == 2 {
         return youtube.Android_Racy.Exchange(v.id, change)
      }
      return youtube.Android_Content.Exchange(v.id, change)
   }
   return youtube.Android.Player(v.id)
}

type video struct {
   address string
   audio string
   height int
   id string
   info bool
   request int
}
