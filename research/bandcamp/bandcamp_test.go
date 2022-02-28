package bandcamp

import (
   "testing"
)

func TestBandcamp(t *testing.T) {
   _, err := newSession("https://hotcasarecords.bandcamp.com/music")
   if err != nil {
      t.Fatal(err)
   }
}
