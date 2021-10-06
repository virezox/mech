package bandcamp

import (
   "testing"
   "time"
)

type test struct {
   in string
   typ string
   id int
}

var tests = []test{
   {"https://schnaussandmunk.bandcamp.com/album/passage-2", "a", 1670971920},
   {"https://schnaussandmunk.bandcamp.com/track/amaris-2", "t", 2809477874},
}

func TestDetail(t *testing.T) {
   Verbose(true)
   for _, test := range tests {
      d, err := TralbumDetail(test.in)
      if err != nil {
         t.Fatal(err)
      }
      if d.Tralbum_Type != test.typ {
         t.Fatal(d.Tralbum_Type)
      }
      if d.Tralbum_ID != test.id {
         t.Fatal(d.Tralbum_ID)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
