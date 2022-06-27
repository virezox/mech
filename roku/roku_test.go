package roku

import (
   "fmt"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const playback_id = "597a64a4a25c5bf6af4a8c7053049a6f"

var client = widevine.Client{Raw: "28339AD78F734520DA24E6E0573D392E"}

func Test_Playback(t *testing.T) {
   site, err := New_Cross_Site()
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
   client.Private_Key, err = os.ReadFile(home + "/mech/private_key.pem")
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
// therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
const id = "105c41ea75775968b670fbb26978ed76"

func Test_Video(t *testing.T) {
   con, err := New_Content(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(con)
   video, err := con.HLS()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
