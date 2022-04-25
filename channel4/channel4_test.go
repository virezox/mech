package channel4

import (
   "fmt"
   "testing"
)

// channel4.com/programmes/frasier/on-demand/18926-001
const frasier = "18926-001"

func TestStream(t *testing.T) {
   video, err := NewVideo(frasier)
   if err != nil {
      t.Fatal(err)
   }
   media, err := video.Widevine().Media()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
