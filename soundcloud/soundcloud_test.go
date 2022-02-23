package soundcloud

import (
   "fmt"
   "testing"
)

const addr =
   "https://soundcloud.com/afterhour-sounds/premiere-ele-bisu-caradamom-coffee"

var ids = []int64{1021056175}

func TestResolve(t *testing.T) {
   track, err := Resolve(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", track)
}

func TestTracks(t *testing.T) {
   tracks, err := Tracks(ids)
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tracks[0].Progressive()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

func TestUser(t *testing.T) {
   tracks, err := UserTracks(692707328)
   if err != nil {
      t.Fatal(err)
   }
   for _, track := range tracks {
      fmt.Printf("%+v\n", track)
   }
}
