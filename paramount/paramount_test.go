package paramount

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = map[test_type]string{
   {episode, dash_cenc}: "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU",
   {episode, stream_pack}: "622520382",
   {movie, stream_pack}: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
}

var client = Client{Raw_Key_ID: "3be8be937c98483184b294173f9152af"}

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

const (
   dash_cenc = iota
   episode
   movie
   stream_pack
)

type test_type struct {
   content_type int
   asset int
}

func TestPreview(t *testing.T) {
   for _, test := range tests {
      preview, err := NewMedia(test).Preview()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", preview)
      time.Sleep(time.Second)
   }
}
