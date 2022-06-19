package roku

import (
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const playback_id = "597a64a4a25c5bf6af4a8c7053049a6f"

var client = Client{RawKeyId: "28339AD78F734520DA24E6E0573D392E"}

func TestPlayback(t *testing.T) {
   site, err := NewCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(playback_id)
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   client.ID, err = os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   client.PrivateKey, err = os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   content, err := play.Content(client)
   if err != nil {
      t.Fatal(err)
   }
   if content.String() != "13d7c7cf295444944b627ef0ad2c1b3c" {
      t.Fatal(content)
   }
}
