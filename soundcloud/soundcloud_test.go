package soundcloud

import (
   "fmt"
   "testing"
)

const addr = "https://soundcloud.com/bluewednesday/murmuration-feat-shopan"

func TestSoundCloud(t *testing.T) {
   track, err := NewTrack(addr)
   if err != nil {
      t.Fatal(err)
   }
   m, err := track.GetMedia()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
