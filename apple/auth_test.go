package apple

import (
   "os"
   "testing"
)

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
   if err := auth.Create(home + "/mech/apple.json"); err != nil {
      t.Fatal(err)
   }
}
