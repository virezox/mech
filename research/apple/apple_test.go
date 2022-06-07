package apple

import (
   "fmt"
   "os"
   "testing"
)

const (
   contentID = "umc.cmc.45cu44369hb2qfuwr3fihnr8e"
   pssh = "data:text/plain;base64,AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="
)

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := OpenAuth(home, "mech/apple.json")
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
   license, err := auth.License(privateKey, clientID, pssh)
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
   key, err := license.Key(env, episode)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
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
