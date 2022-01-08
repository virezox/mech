package main

import (
   "fmt"
   "os"
   "time"
   gp "github.com/89z/googleplay"
)

var apps = []string{
   "bbc.mobile.news.ww",
   "com.amazon.mp3",
   "com.aspiro.tidal",
   "com.bandcamp.android",
   "com.clearchannel.iheartradio.controller",
   "com.google.android.youtube",
   "com.instagram.android",
   "com.nbcuni.nbc",
   "com.pandora.android",
   "com.pbs.video",
   "com.qobuz.music",
   "com.reddit.frontpage",
   "com.rhapsody",
   "com.soundcloud.android",
   "com.spotify.music",
   "com.twitter.android",
   "com.vimeo.android.videoapp",
   "com.zhiliaoapp.musically",
   "deezer.android.app",
}

func main() {
   auth, err := newAuth()
   if err != nil {
      panic(err)
   }
   dev, err := newDevice()
   if err != nil {
      panic(err)
   }
   for _, app := range apps {
      det, err := auth.Details(dev, app)
      if err != nil {
         panic(err)
      }
      fmt.Println(det)
      time.Sleep(time.Second)
   }
}

func newDevice() (*gp.Device, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   file, err := os.Open(cache + "/googleplay/device.json")
   if err != nil {
      return nil, err
   }
   defer file.Close()
   return gp.ReadDevice(file)
}

func newAuth() (*gp.Auth, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   file, err := os.Open(cache + "/googleplay/token.json")
   if err != nil {
      return nil, err
   }
   defer file.Close()
   tok, err := gp.ReadToken(file)
   if err != nil {
      return nil, err
   }
   return tok.Auth()
}

