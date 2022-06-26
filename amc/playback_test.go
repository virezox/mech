package amc

import (
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
const nID = 1011152

var client = widevine.Client{Raw: "c0e598b247fa443590299d5ef47da32c"}

func Test_Playback(t *testing.T) {
   home, err := os.UserHomeDir()
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
   client.ID, err = os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   client.Private_Key, err = os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   content, err := play.Content(client)
   if err != nil {
      t.Fatal(err)
   }
   if content.String() != "a66a5603545ad206c1a78e160a6710b1" {
      t.Fatal(content)
   }
}
