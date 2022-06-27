package roku

import (
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const (
   key = "13d7c7cf295444944b627ef0ad2c1b3c"
   playback_ID = "597a64a4a25c5bf6af4a8c7053049a6f"
   raw_key_ID = "28339AD78F734520DA24E6E0573D392E"
)

func Test_Post(t *testing.T) {
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
   key_ID, err := widevine.Key_ID(raw_key_ID)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      t.Fatal(err)
   }
   site, err := New_Cross_Site()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(playback_ID)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(play)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != key {
      t.Fatal(keys)
   }
}
