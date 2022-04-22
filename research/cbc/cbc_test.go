package cbc

import (
   "fmt"
   "testing"
)

func TestOTT(t *testing.T) {
   login := Login{Access_Token: accessToken}
   web, err := login.WebToken()
   if err != nil {
      t.Fatal(err)
   }
   top, err := web.OverTheTop()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", top)
}

func TestLogin(t *testing.T) {
   login, err := NewLogin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login)
}
