package cbc

import (
   "fmt"
   "os"
   "testing"
)

const downton = "downton-abbey/s01e05"

func TestMedia(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := OpenProfile(cache, "mech/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   asset, err := NewAsset(downton)
   if err != nil {
      t.Fatal(err)
   }
   media, err := profile.Media(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}

func TestProfile(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := NewLogin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   web, err := login.WebToken()
   if err != nil {
      t.Fatal(err)
   }
   top, err := web.OverTheTop()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := top.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := profile.Create(cache, "mech/cbc.json"); err != nil {
      t.Fatal(err)
   }
}
