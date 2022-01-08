package pandora

import (
   "fmt"
   "os"
   "testing"
   "time"
)

// Note that you cannot get RADIO tracks with this method.
var pandoraIDs = []string{
   // pandora.com/artist/the-black-dog/radio-scarecrow/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV
   "TR:1168891",
   // pandora.com/artist/the-black-dog/music-for-real-airports/m-1/TRnJq99pmqt72Zc
   "TR:1616369",
   // pandora.com/artist/jessy-lanza/pull-my-hair-back/strange-emotion/TRkbfrm9rfpZZbq
   "TR:2314875",
}

func TestDecode(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Open(cache + "/mech/pandora.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   user := new(UserLogin)
   if err := user.Decode(file); err != nil {
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

func TestEncode(t *testing.T) {
   part, err := NewPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   user, err := part.UserLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      t.Fatal(tLen)
   }
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   cache += "/mech"
   os.Mkdir(cache, os.ModeDir)
   file, err := os.Create(cache + "/pandora.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   if err := user.Encode(file); err != nil {
      t.Fatal(err)
   }
}
