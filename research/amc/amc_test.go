package amc

import (
   "os"
   "testing"
)

// amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
const nid = 1011152

func TestAMC(t *testing.T) {
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(email, password); err != nil {
      t.Fatal(err)
   }
   res, err := auth.Playback(nid)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
