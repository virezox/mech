package resolve

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "https://schnaussandmunk.bandcamp.com/track/amaris-2",
   "https://schnaussandmunk.bandcamp.com/album/passage-2",
}

func TestResolve(t *testing.T) {
   for _, test := range tests {
      d, err := newDetails(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", d)
      time.Sleep(99 * time.Millisecond)
   }
}
