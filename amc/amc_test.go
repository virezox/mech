package amc

import (
   "encoding/hex"
   "os"
   "testing"
)

const (
   // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
   nid = 1011152
   rawKID = "c0e598b247fa443590299d5ef47da32c"
)

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
   privateKey, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   kID, err := hex.DecodeString(rawKID)
   if err != nil {
      t.Fatal(err)
   }
   key, err := play.Key(privateKey, clientID, kID)
   if err != nil {
      t.Fatal(err)
   }
   if hex.EncodeToString(key) != "a66a5603545ad206c1a78e160a6710b1" {
      t.Fatal(key)
   }
}

func TestRefresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := OpenAuth(home, "mech/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   if err := auth.Create(home, "mech/amc.json"); err != nil {
      t.Fatal(err)
   }
}

func TestLogin(t *testing.T) {
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
