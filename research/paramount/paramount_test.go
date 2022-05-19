package paramount

import (
   "encoding/hex"
   "github.com/89z/format/dash"
   "github.com/89z/mech/research/widevine"
   "os"
   "testing"
)

const contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"

func TestParamount(t *testing.T) {
   privateKey, err := os.ReadFile("../widevine/ignore/device_private_key")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile("../widevine/ignore/device_client_id_blob")
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Open("ignore.mpd")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   adas, err := dash.NewAdaptationSet(file)
   if err != nil {
      t.Fatal(err)
   }
   kID, err := adas.Protection().KID()
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
   keys, err := mod.Keys(sess.URL, sess.Header())
   if err != nil {
      t.Fatal(err)
   }
   var pass bool
   for _, key := range keys {
      if hex.EncodeToString(key.Key) == "44f12639c9c4a5a432338aca92e38920" {
         pass = true
      }
   }
   if !pass {
      t.Fatal(keys)
   }
}
