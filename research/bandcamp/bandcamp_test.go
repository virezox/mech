package bandcamp

import (
   "testing"
)

func TestBandcamp(t *testing.T) {
   err := NewDataTralbum("https://hotcasarecords.bandcamp.com/music")
   if err != nil {
      t.Fatal(err)
   }
}
