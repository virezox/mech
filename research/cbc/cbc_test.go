package cbc

import (
   "fmt"
   "testing"
)

const downton = "https://gem.cbc.ca/media/downton-abbey/s01e05"

func TestMedia(t *testing.T) {
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
   asset, err := NewAsset(downton)
   if err != nil {
      t.Fatal(err)
   }
   media, err := pro.Media(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}

func TestLogin(t *testing.T) {
   login, err := NewLogin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login)
}
