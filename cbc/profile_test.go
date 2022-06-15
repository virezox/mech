package cbc

import (
   "os"
   "testing"
)

func TestProfile(t *testing.T) {
   home, err := os.UserHomeDir()
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
   if err := profile.Create(home + "/mech/cbc.json"); err != nil {
      t.Fatal(err)
   }
}
