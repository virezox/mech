package bbcamerica

import (
   "github.com/89z/format/dash"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// bbcamerica.com/shows/killing-eve/episodes/season-4-just-dunk-me--1052529
const nid = 1052529

func TestUnauth(t *testing.T) {
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
   file, err := os.Open("ignore.mpd")
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
   auth, err := NewUnauth()
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback(nid)
   if err != nil {
      t.Fatal(err)
   }
   addr := play.DASH().Key_Systems.Widevine.License_URL
   widevine.LogLevel = 1
   if _, err := mod.Post(addr, auth.Header()); err != nil {
      t.Fatal(err)
   }
}
