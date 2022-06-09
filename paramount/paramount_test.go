package paramount

import (
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

const rawKID = "3be8be937c98483184b294173f9152af"

var tests = map[testType]string{
   {episode, dashCenc}: "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU",
   {episode, streamPack}: "622520382",
   {movie, streamPack}: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
}

func TestSession(t *testing.T) {
   sess, err := NewSession(tests[testType{episode, dashCenc}])
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
   kID, err := hex.DecodeString(rawKID)
   if err != nil {
      t.Fatal(err)
   }
   key, err := sess.Key(privateKey, clientID, kID)
   if err != nil {
      t.Fatal(err)
   }
   if hex.EncodeToString(key) != "44f12639c9c4a5a432338aca92e38920" {
      t.Fatal(key)
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
