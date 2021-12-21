package pandora

import (
   "fmt"
   "os"
   "testing"
)

const pandoraID = "TR:1168891"

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
   info, err := user.PlaybackInfo(pandoraID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", info)
}

func TestEncode(t *testing.T) {
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
      t.Fatal(err)
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
