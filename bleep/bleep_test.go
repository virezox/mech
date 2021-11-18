package bleep

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

const addr = "https://bleep.com/release/8728-four-tet-pause"

func TestMeta(t *testing.T) {
   mech.Verbose = true
   met, err := NewMeta(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", met)
   date, err := met.Release_Date.Parse()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(date)
}

func TestResolve(t *testing.T) {
   mech.Verbose = true
   tracks, err := Release(addr)
   if err != nil {
      t.Fatal(err)
   }
   addr, err := tracks[0].Resolve()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(addr)
}
