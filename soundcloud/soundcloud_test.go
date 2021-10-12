package soundcloud

import (
   "fmt"
   "testing"
)

const (
   addr = "https://soundcloud.com/afterhour-sounds/premiere-ele-bisu-caradamom-coffee"
   ids = "1021056175"
)

func TestAlternate(t *testing.T) {
   Verbose(true)
   a, err := Oembed(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}

func TestResolve(t *testing.T) {
   Verbose(true)
   r, err := Resolve(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", r)
}

func TestTracks(t *testing.T) {
   Verbose(true)
   tracks, err := Tracks(ids)
   if err != nil {
      t.Fatal(err)
   }
   m, err := tracks[0].GetMedia()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
