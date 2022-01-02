package bleep

import (
   "fmt"
   "testing"
)

const releaseID = 8728

func TestMeta(t *testing.T) {
   meta, err := NewMeta(releaseID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", meta)
   date, err := meta.ReleaseDate()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(date)
}

func TestResolve(t *testing.T) {
   LogLevel = 1
   tracks, err := Release(releaseID)
   if err != nil {
      t.Fatal(err)
   }
   addr, err := tracks[0].Resolve()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(addr)
}
