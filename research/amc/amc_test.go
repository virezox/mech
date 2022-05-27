package amc

import (
   "fmt"
   "github.com/89z/format/dash"
   "github.com/89z/mech/widevine"
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
   privateKey, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Open("amc.mpd")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   period, err := dash.NewPeriod(file)
   if err != nil {
      t.Fatal(err)
   }
   kID, err := period.Protection().KID()
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(privateKey, clientID, kID)
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
   addr := play.DASH().Key_Systems.Widevine.License_URL
   keys, err := mod.Post(addr, play.Header())
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", keys)
   return
   if keys.Content().String() != "680a46ebd6cf2b9a6a0b05a24dcf944a" {
      t.Fatal(keys)
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
