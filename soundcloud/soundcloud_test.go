package soundcloud

import (
   "fmt"
   "testing"
)

const addr = "https://soundcloud.com/bluewednesday/murmuration-feat-shopan"

func TestSoundCloud(t *testing.T) {
   id, err := ClientID()
   if err != nil {
      t.Fatal(err)
   }
   track, err := NewTrack(id, addr)
   if err != nil {
      t.Fatal(err)
   }
   for _, code := range track.Media.Transcodings {
      fmt.Printf("%+v\n", code)
   }
}
