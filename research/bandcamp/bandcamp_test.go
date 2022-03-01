package bandcamp

import (
   "testing"
   "time"
)

type testType struct {
   typ, addr string
}

var tests = []testType{
   {"a", "https://schnaussandmunk.bandcamp.com/album/passage-2"},
   {"i", "https://schnaussandmunk.bandcamp.com/music"},
   {"t", "https://schnaussandmunk.bandcamp.com/track/amaris-2"},
}

func TestToken(t *testing.T) {
   for _, test := range tests {
      tok, err := NewToken(test.addr)
      if err != nil {
         t.Fatal(err)
      }
      if tok.Type != test.typ {
         t.Fatal(tok)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
