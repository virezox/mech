package amc

import (
   "fmt"
   "os"
   "testing"
)

// amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
const nid = 1011152

func TestPlayback(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := OpenAuth(home, "mech/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback(nid)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}

func TestCreate(t *testing.T) {
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(email, password); err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Create(home, "mech/amc.json"); err != nil {
      t.Fatal(err)
   }
}
