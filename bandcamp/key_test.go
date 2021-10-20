package bandcamp

import (
   "fmt"
   "testing"
)

const track = "https://schnaussandmunk.bandcamp.com/track/amaris-2"

func TestTrack(t *testing.T) {
   Verbose(true)
   iURL, err := NewInfoURL(track)
   if err != nil {
      t.Fatal(err)
   }
   track, err := NewInfoTrack(iURL.Track_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", track)
}
