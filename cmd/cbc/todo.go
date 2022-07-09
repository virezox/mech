package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/cbc"
   "github.com/89z/rosso/hls"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
   "strings"
)

func (f flags) master() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   profile, err := cbc.Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      return err
   }
   asset, err := cbc.New_Asset(f.id)
   if err != nil {
      return err
   }
   asset_media, err := profile.Media(asset)
   if err != nil {
      return err
   }
   master, err := f.Master(*asset_media.URL, asset.AppleContentId)
   if err != nil {
      return err
   }
   // audio
   media := master.Media.Filter(func(m hls.Media) bool {
      return m.Type == "AUDIO"
   })
   medium := media.Reduce(func(carry, item hls.Media) bool {
      return item.Name == f.audio_name
   })
   // video
   streams := master.Stream.Filter(func(s hls.Stream) bool {
      return strings.HasPrefix(s.Codecs, "avc1.")
   })
   stream := streams.Reduce(func(carry, item hls.Stream) bool {
      distance := hls.Bandwidth(f.video_bandwidth)
      return distance(item) < distance(carry)
   })
}
