package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/googleplay"
   "os"
   "sort"
   "time"
)

var apps = []string{
   "bbc.mobile.news.ww",
   "com.amazon.mp3",
   "com.aspiro.tidal",
   "com.bandcamp.android",
   "com.cbs.app",
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
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   tok, err := googleplay.OpenToken(cache + "/googleplay/token.json")
   if err != nil {
      panic(err)
   }
   auth, err := tok.Auth()
   if err != nil {
      panic(err)
   }
   dev, err := googleplay.OpenDevice(cache + "/googleplay/device.json")
   if err != nil {
      panic(err)
   }
   var dets []*googleplay.Details
   for _, app := range apps {
      det, err := auth.Details(dev, app)
      if err != nil {
         panic(err)
      }
      dets = append(dets, det)
      time.Sleep(99 * time.Millisecond)
   }
   sort.Slice(dets, func(a, b int) bool {
      return dets[b].NumDownloads < dets[a].NumDownloads
   })
   for _, det := range dets {
      fmt.Println(format.Number.GetUint64(det.NumDownloads), det.Title)
   }
}
