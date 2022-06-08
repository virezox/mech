package roku

import (
   "encoding/hex"
   "github.com/89z/format/dash"
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const playbackID = "597a64a4a25c5bf6af4a8c7053049a6f"

func TestWidevine(t *testing.T) {
   site, err := NewCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(playbackID)
   if err != nil {
      t.Fatal(err)
   }
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
   key, err := play.Key(privateKey, clientID, kID)
   if err != nil {
      t.Fatal(err)
   }
   if hex.EncodeToString(key) != "13d7c7cf295444944b627ef0ad2c1b3c" {
      t.Fatal(key)
   }
}
