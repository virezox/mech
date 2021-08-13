package soundcloud

import (
   "fmt"
   "testing"
)

const addr = "https://soundcloud.com/bluewednesday/murmuration-feat-shopan"

func TestSoundCloud(t *testing.T) {
   err := WriteClientID()
   if err != nil {
      t.Fatal(err)
   }
   id, err := ReadClientID()
   if err != nil {
      t.Fatal(err)
   }
   track, err := NewTrack(id, addr)
   if err != nil {
      t.Fatal(err)
   }
   m, err := track.GetMedia(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
