package pandora

import (
   "fmt"
   "os"
   "testing"
   "time"
)

const addr =
   "https://pandora.com/artist/the-black-dog/radio-scarecrow" +
   "/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV"

// Note that you cannot get RADIO tracks with this method.
var pandoraIDs = []string{
   // pandora.com/artist/the-black-dog/radio-scarecrow/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV
   "TR:1168891",
   // pandora.com/artist/the-black-dog/music-for-real-airports/m-1/TRnJq99pmqt72Zc
   "TR:1616369",
   // pandora.com/artist/jessy-lanza/pull-my-hair-back/strange-emotion/TRkbfrm9rfpZZbq
   "TR:2314875",
}

func TestMusic(t *testing.T) {
   rec, err := NewMusicRecording(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", rec)
}

func TestOpen(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   user, err := OpenUserLogin(cache + "/mech/pandora.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := user.ValueExchange(); err != nil {
      t.Fatal(err)
   }
   for _, id := range pandoraIDs {
      info, err := user.PlaybackInfo(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("Stat:%v Result:%+v\n", info.Stat, info.Result)
      time.Sleep(time.Second)
   }
}
