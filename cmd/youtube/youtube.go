package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/89z/rosso/os"
)

func (self flags) download() error {
   play, err := self.player()
   if err != nil {
      return err
   }
   forms := play.StreamingData.AdaptiveFormats
   if self.info {
      text, err := play.MarshalText()
      if err != nil {
         return err
      }
      os.Stdout.Write(text)
   } else {
      fmt.Println(play.PlayabilityStatus)
      if self.audio != "" {
         form, ok := forms.Audio(self.audio)
         if ok {
            err := download(form, play.Base())
            if err != nil {
               return err
            }
         }
      }
      if self.height >= 1 {
         form, ok := forms.Video(self.height)
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

func (self flags) player() (*youtube.Player, error) {
   var req youtube.Request
   if self.request == 0 {
      req = youtube.Android()
   } else if self.request == 1 {
      req = youtube.Android_Embed()
   } else {
      if self.request == 2 {
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
   return req.Player(self.video_ID)
}
