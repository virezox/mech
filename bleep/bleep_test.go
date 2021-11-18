package bleep

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

func TestBleep(t *testing.T) {
   id, err := ReleaseID("https://bleep.com/release/8728-four-tet-pause")
   if err != nil {
      t.Fatal(err)
   }
   mech.Verbose = true
   tracks, err := Release(id)
   if err != nil {
      t.Fatal(err)
   }
   addr, err := tracks[0].Resolve()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(addr)
}
