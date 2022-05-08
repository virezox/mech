package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   //"https://schnaussandmunk.bandcamp.com/track/amaris-2",
   //"https://schnaussandmunk.bandcamp.com/album/passage-2",
   "https://schnaussandmunk.bandcamp.com/music",
}

func TestData(t *testing.T) {
   for _, test := range tests {
      data, err := NewData(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", data)
      time.Sleep(99 * time.Millisecond)
   }
}
