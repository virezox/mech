package bandcamp

import (
   "fmt"
   "testing"
)

const track = "https://schnaussandmunk.bandcamp.com/track/amaris-2"

func TestTrack(t *testing.T) {
   Verbose(true)
   ai, err := NewUrlInfo(track)
   if err != nil {
      t.Fatal(err)
   }
   ti, err := NewTrackInfo(ai.Track_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ti)
}
