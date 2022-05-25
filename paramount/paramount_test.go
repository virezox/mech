package paramount

import (
   "fmt"
   "github.com/89z/format/dash"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
   "time"
)

var guids = []string{
   // paramountplus.com/movies/building-star-trek/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
   "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_",
   // paramountplus.com/shows/melrose_place/video/622520382/melrose-place-pilot
   "622520382",
}

func TestParamount(t *testing.T) {
   for _, guid := range guids {
      preview, err := NewMedia(guid).Preview()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", preview)
      time.Sleep(time.Second)
   }
}

const contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"

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
   sess, err := NewSession(contentID)
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
