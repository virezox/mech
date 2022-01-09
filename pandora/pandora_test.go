package pandora

import (
   "bytes"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

const helloEnc = "7be654d97cc31582815d713a9d0c64ab"

var helloDec = []byte("hello world")

func TestLogin(t *testing.T) {
   part, err := NewPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   tLen := len(part.Result.PartnerAuthToken)
   if tLen != 34 {
      t.Fatal(tLen)
   }
   user, err := part.UserLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   if tLen := len(user.Result.UserAuthToken); tLen != 58 {
      t.Fatal(tLen)
   }
}

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(helloEnc)
   if err != nil {
      t.Fatal(err)
   }
   dec, err := Decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(dec, helloDec) {
      t.Fatal(dec)
   }
}

// Note that you cannot get RADIO tracks with this method.
var pandoraIDs = []string{
   // pandora.com/artist/the-black-dog/radio-scarecrow/train-by-the-autobahn-part-1/TRddpp5JJ2hqnVV
   "TR:1168891",
   // pandora.com/artist/the-black-dog/music-for-real-airports/m-1/TRnJq99pmqt72Zc
   "TR:1616369",
   // pandora.com/artist/jessy-lanza/pull-my-hair-back/strange-emotion/TRkbfrm9rfpZZbq
   "TR:2314875",
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

func TestCreate(t *testing.T) {
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
   if err := user.Create(cache + "/mech/pandora.json"); err != nil {
      t.Fatal(err)
   }
}
