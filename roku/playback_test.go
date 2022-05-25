package roku

import (
   "github.com/89z/format/dash"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const playbackID = "597a64a4a25c5bf6af4a8c7053049a6f"

func TestWidevine(t *testing.T) {
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
   site, err := NewCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(playbackID)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(play.DRM.Widevine.LicenseServer, nil)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != "13d7c7cf295444944b627ef0ad2c1b3c" {
      t.Fatal(keys)
   }
}
