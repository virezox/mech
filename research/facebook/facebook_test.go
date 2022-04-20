package facebook

import (
   "fmt"
   "testing"
)

func TestFacebook(t *testing.T) {
   login, err := NewLogin()
   if err != nil {
      t.Fatal(err)
   }
   reg, err := login.Regular(email, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", reg)
}
