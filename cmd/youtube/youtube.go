package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

func (f flags) download() error {
   play, err := f.player()
   if err != nil {
      return err
   }
   forms := play.StreamingData.AdaptiveFormats
   if f.info {
      text, err := play.MarshalText()
      if err != nil {
         return err
      }
      os.Stdout.Write(text)
   } else {
      fmt.Println(play.PlayabilityStatus)
      if f.audio != "" {
         form, ok := forms.Audio(f.audio)
         if ok {
            err := download(form, play.Base())
            if err != nil {
               return err
            }
         }
      }
      if f.height >= 1 {
         form, ok := forms.Video(f.height)
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

func refresh() error {
   auth, err := youtube.New_OAuth()
   if err != nil {
      return err
   }
   fmt.Println(auth)
   fmt.Scanln()
   head, err := auth.Header()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return head.Create(home + "/mech/youtube.json")
}

func access() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   head, err := youtube.Open_Header(home + "/mech/youtube.json")
   if err != nil {
      return err
   }
   if err := head.Refresh(); err != nil {
      return err
   }
   return head.Create(home + "/mech/youtube.json")
}

func download(form *youtube.Format, base string) error {
   ext, err := form.Ext()
   if err != nil {
      return err
   }
   file, err := os.Create(base + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   return form.Encode(file)
}

func (f flags) player() (*youtube.Player, error) {
   var req youtube.Request
   if f.request == 0 {
      req = youtube.Android()
   } else if f.request == 1 {
      req = youtube.Android_Embed()
   } else {
      if f.request == 2 {
         req = youtube.Android_Racy()
      } else {
         req = youtube.Android_Content()
      }
      home, err := os.UserHomeDir()
      if err != nil {
         return nil, err
      }
      req.Header, err = youtube.Open_Header(home + "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
   }
   return req.Player(f.video_ID)
}
