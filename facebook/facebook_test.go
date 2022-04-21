package facebook

import (
   "fmt"
   "testing"
)

const addr =
   "https://www.facebook.com/FromTheBasementPage/videos/309868367063220"

func TestMeta(t *testing.T) {
   meta, err := NewMeta(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", meta)
}

const id = 444624393796648

func TestRegular(t *testing.T) {
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
