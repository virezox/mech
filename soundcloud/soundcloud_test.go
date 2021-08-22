package soundcloud

import (
   "fmt"
   "testing"
)

const (
   addr = "https://soundcloud.com/pdis_inpartmaint/harold-budd-perhaps-moss"
   id = "103650107"
)

func TestID(t *testing.T) {
   tracks, err := Tracks(id)
   if err != nil {
      t.Fatal(err)
   }
   m, err := tracks[0].GetMedia()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}

func TestURL(t *testing.T) {
   track, err := Resolve(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", track)
}
