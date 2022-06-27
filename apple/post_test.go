package apple

import (
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// tv.apple.com/us/episode/biscuits/umc.cmc.45cu44369hb2qfuwr3fihnr8e
const (
   content_ID = "umc.cmc.45cu44369hb2qfuwr3fihnr8e"
   key = "22bdb0063805260307ee5045c0f3835a"
   pssh = "data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="
)

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   key_ID, err := widevine.PSSH_Key_ID(pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Open_Auth(home + "/mech/apple.json")
   if err != nil {
      t.Fatal(err)
   }
   env, err := New_Environment()
   if err != nil {
      t.Fatal(err)
   }
   episode, err := New_Episode(content_ID)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(Poster{auth, env, episode, pssh})
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != key {
      t.Fatal(keys)
   }
}
