package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const artID = 3809045440

func TestImage(t *testing.T) {
   for _, img := range Images {
      addr := img.URL(artID)
      fmt.Println(addr)
   }
}

var tests = []string{
   "https://schnaussandmunk.bandcamp.com",
   "https://schnaussandmunk.bandcamp.com/album/passage-2",
   "https://schnaussandmunk.bandcamp.com/music",
   "https://schnaussandmunk.bandcamp.com/track/amaris-2",
}

func TestParam(t *testing.T) {
   for _, test := range tests {
      param, err := NewParams(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", param)
      time.Sleep(99 * time.Millisecond)
   }
}
