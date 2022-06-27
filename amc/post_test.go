package amc

import (
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
const nID = 1011152

func Test_Playback(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   key_ID, err := widevine.Key_ID("c0e598b247fa443590299d5ef47da32c")
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Open_Auth(home + "/mech/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback(nID)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(play)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != "a66a5603545ad206c1a78e160a6710b1" {
      t.Fatal(keys)
   }
}
