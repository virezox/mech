package apple

import (
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

// tv.apple.com/us/episode/biscuits/umc.cmc.45cu44369hb2qfuwr3fihnr8e
const contentID = "umc.cmc.45cu44369hb2qfuwr3fihnr8e"

var client = widevine.Client{RawPSSH: "data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="}

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := OpenAuth(home, "mech/apple.json")
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
   env, err := NewEnvironment()
   if err != nil {
      t.Fatal(err)
   }
   episode, err := NewEpisode(contentID)
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
func TestCreate(t *testing.T) {
   con, err := NewConfig()
   if err != nil {
      t.Fatal(err)
   }
   sign, err := con.Signin(email, password)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := sign.Auth()
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Create(home, "mech/apple.json"); err != nil {
      t.Fatal(err)
   }
}
