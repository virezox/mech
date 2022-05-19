package roku

import (
   "github.com/89z/format/dash"
   "github.com/89z/mech/roku"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const playbackID = "597a64a4a25c5bf6af4a8c7053049a6f"

func TestPlayback(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   privateKey, err := os.ReadFile(cache + "/mech/device_private_key")
   if err != nil {
      t.Fatal(err)
   }
   clientID, err := os.ReadFile(cache + "/mech/device_client_id_blob")
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
   site, err := roku.NewCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(playbackID)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(play.DRM.Widevine.LicenseServer, nil)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != "13d7c7cf295444944b627ef0ad2c1b3c" {
      t.Fatal(keys)
   }
}
