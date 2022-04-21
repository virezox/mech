package facebook

import (
   "fmt"
   "testing"
)

const (
   anon = 309868367063220
   auth = 444624393796648
)

func TestMeta(t *testing.T) {
   meta, err := NewMeta(anon)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", meta)
}

func TestRegular(t *testing.T) {
   login, err := NewLogin()
   if err != nil {
      t.Fatal(err)
   }
   reg, err := login.Regular(email, password)
   if err != nil {
      t.Fatal(err)
   }
   vid, err := reg.Video(auth)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
