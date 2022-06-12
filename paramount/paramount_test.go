package paramount

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = map[testType]string{
   {episode, dashCenc}: "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU",
   {episode, streamPack}: "622520382",
   {movie, streamPack}: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
}

var client = Client{RawKeyID: "3be8be937c98483184b294173f9152af"}

func TestSession(t *testing.T) {
   sess, err := NewSession(tests[testType{episode, dashCenc}])
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
   content, err := sess.Content(client)
   if err != nil {
      t.Fatal(err)
   }
   if content.String() != "44f12639c9c4a5a432338aca92e38920" {
      t.Fatal(content)
   }
}

const (
   dashCenc = iota
   episode
   movie
   streamPack
)

type testType struct {
   contentType int
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
