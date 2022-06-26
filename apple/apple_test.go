package apple

import (
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// tv.apple.com/us/episode/biscuits/umc.cmc.45cu44369hb2qfuwr3fihnr8e
const content_ID = "umc.cmc.45cu44369hb2qfuwr3fihnr8e"

var client = widevine.Client{Raw: "data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="}

func Test_License(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Open_Auth(home + "/mech/apple.json")
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
   env, err := New_Environment()
   if err != nil {
      t.Fatal(err)
   }
   episode, err := New_Episode(content_ID)
   if err != nil {
      t.Fatal(err)
   }
   request, err := auth.Request(client)
   if err != nil {
      t.Fatal(err)
   }
   license, err := request.License(env, episode)
   if err != nil {
      t.Fatal(err)
   }
   content, err := license.Content()
   if err != nil {
      t.Fatal(err)
   }
   if content.String() != "22bdb0063805260307ee5045c0f3835a" {
      t.Fatal(content)
   }
}
