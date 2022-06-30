package apple

import (
   "fmt"
   "os"
   "testing"
)

func Test_Asset(t *testing.T) {
   episode, err := New_Episode(content_ID)
   if err != nil {
      t.Fatal(err)
   }
   asset := episode.Asset()
   fmt.Printf("%+v\n", asset)
}

func Test_Create(t *testing.T) {
   con, err := New_Config()
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
   if err := auth.Create(home + "/mech/apple.json"); err != nil {
      t.Fatal(err)
   }
}
