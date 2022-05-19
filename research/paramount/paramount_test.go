package paramount

import (
   "github.com/89z/format/dash"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

const contentID = "eyT_RYkqNuH_6ZYrepLtxkiPO1HA7dIU"

func TestParamount(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   privateKey, err := os.ReadFile(cache + "/mech/device_private_key")
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
   mod, err := widevine.NewModule(privateKey, mes.Marshal(), kID)
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
