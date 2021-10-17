package bandcamp

import (
   "fmt"
   "testing"
)

const track = "https://schnaussandmunk.bandcamp.com/track/amaris-2"

func TestTrack(t *testing.T) {
   Verbose(true)
   inf, err := NewInfo(track)
   if err != nil {
      t.Fatal(err)
   }
   tra, err := NewTrack(inf.Track_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", tra)
}
