package paramount

import (
   "encoding/hex"
   "github.com/89z/mech/research/widevine"
   "os"
   "testing"
)

const contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"

func TestParamount(t *testing.T) {
   media, err := NewMedia("ignore.mpd")
   if err != nil {
      t.Fatal(err)
   }
   privateKey, err := os.ReadFile("../widevine/ignore/device_private_key")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile("../widevine/ignore/device_client_id_blob")
   if err != nil {
      t.Fatal(err)
   }
   keyID, err := hex.DecodeString(media.KID())
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(privateKey, clientID, keyID)
   if err != nil {
      t.Fatal(err)
   }
   session, err := NewSession()
   if err != nil {
      t.Fatal(err)
   }
   keys, err := session.Keys(contentID, mod)
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
