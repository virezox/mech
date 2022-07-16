package main

import (
   "github.com/89z/mech/cbc"
   "github.com/89z/rosso/hls"
   "os"
)

func (self flags) download() error {
   master, err := self.master()
   if err != nil {
      return err
   }
   media := master.Media.Filter(func(m hls.Medium) bool {
      return m.Type == "AUDIO"
   })
   index := media.Index(func(a, b hls.Medium) bool {
      return b.Name == self.name
   })
   if err := self.HLS_Media(media, index); err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   return self.HLS_Streams(streams, streams.Bandwidth(self.bandwidth))
}

func (self *flags) master() (*hls.Master, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   profile, err := cbc.Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      return nil, err
   }
   asset, err := cbc.New_Asset(self.id)
   if err != nil {
      return nil, err
   }
   media, err := profile.Media(asset)
   if err != nil {
      return nil, err
   }
   self.Base = asset.AppleContentId
   return self.HLS(*media.URL)
}

func (self flags) profile() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   login, err := cbc.New_Login(self.email, self.password)
   if err != nil {
      return err
   }
   web, err := login.Web_Token()
   if err != nil {
      return err
   }
   top, err := web.Over_The_Top()
   if err != nil {
      return err
   }
   profile, err := top.Profile()
   if err != nil {
      return err
   }
   return profile.Create(home + "/mech/cbc.json")
}
