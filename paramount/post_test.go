package paramount

import (
   "fmt"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
   "time"
)

var client = widevine.Client{Raw: "3be8be937c98483184b294173f9152af"}

func Test_Session(t *testing.T) {
   sess, err := New_Session(tests[test_type{episode, dash_cenc}])
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
   content, err := sess.Content(client)
   if err != nil {
      t.Fatal(err)
   }
   if content.String() != "44f12639c9c4a5a432338aca92e38920" {
      t.Fatal(content)
   }
}
