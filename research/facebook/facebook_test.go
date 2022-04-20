package facebook

import (
   "fmt"
   "testing"
)

const id = 444624393796648

func TestFacebook(t *testing.T) {
   login, err := NewLogin()
   if err != nil {
      t.Fatal(err)
   }
   reg, err := login.Regular(email, password)
   if err != nil {
      t.Fatal(err)
   }
   vid, err := reg.Video(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
