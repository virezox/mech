package apple

import (
   "bytes"
   "encoding/base64"
   "fmt"
   "github.com/89z/mech/widevine"
   "os"
   "testing"
)

const contentID = "umc.cmc.45cu44369hb2qfuwr3fihnr8e"

func TestLicense(t *testing.T) {
   episode, err := NewEpisode(contentID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", episode)
   env, err := NewEnvironment()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", env)
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
   kID, err := base64.StdEncoding.DecodeString("AAAAABaC7ytjMCAgICAgIA==")
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(privateKey, clientID, kID)
   if err != nil {
      t.Fatal(err)
   }
   body, err := base64.StdEncoding.DecodeString(license)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Keys(bytes.NewReader(body))
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", keys)
}
