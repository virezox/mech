package paramount

import (
   "fmt"
   "github.com/89z/format/dash"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
   "time"
)

var tests = map[testType]string{
   {episode, dashCenc}: "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU",
   {episode, streamPack}: "622520382",
   {movie, streamPack}: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
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

func TestSession(t *testing.T) {
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
   sess, err := NewSession(tests[testType{episode, dashCenc}])
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(sess.URL, sess.Header())
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != "44f12639c9c4a5a432338aca92e38920" {
      t.Fatal(keys)
   }
}
