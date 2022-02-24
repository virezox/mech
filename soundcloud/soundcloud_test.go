package soundcloud

import (
   "fmt"
   "testing"
)

const addr =
   "https://soundcloud.com/afterhour-sounds/premiere-ele-bisu-caradamom-coffee"

const (
   trackID = 1021056175
   userID = 692707328
)

func TestTrack(t *testing.T) {
   track, err := NewTrack(trackID)
   if err != nil {
      t.Fatal(err)
   }
   pro, err := track.Progressive()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

func TestResolve(t *testing.T) {
   track, err := Resolve(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", track)
}

func TestUser(t *testing.T) {
   tracks, err := UserTracks(userID)
   if err != nil {
      t.Fatal(err)
   }
   for _, track := range tracks {
      fmt.Printf("%+v\n", track)
   }
}
