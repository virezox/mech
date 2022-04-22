package cbc

import (
   "fmt"
   "os"
   "testing"
)

const downton = "https://gem.cbc.ca/media/downton-abbey/s01e05"

func TestMedia(t *testing.T) {
   asset, err := NewAsset(downton)
   if err != nil {
      t.Fatal(err)
   }
   res, err := asset.Media()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

func TestProfile(t *testing.T) {
   login := Login{Access_Token: accessToken}
   web, err := login.WebToken()
   if err != nil {
      t.Fatal(err)
   }
   top, err := web.OverTheTop()
   if err != nil {
      t.Fatal(err)
   }
   pro, err := top.Profile()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

func TestLogin(t *testing.T) {
   login, err := NewLogin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login)
}
