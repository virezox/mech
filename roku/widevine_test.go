package roku

import (
   "fmt"
   "testing"
)

// therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
const playbackID = "597a64a4a25c5bf6af4a8c7053049a6f"

func TestPlayback(t *testing.T) {
   site, err := NewCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(playbackID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
