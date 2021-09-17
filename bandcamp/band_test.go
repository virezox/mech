package bandcamp

import (
   "fmt"
   "testing"
)

func TestBand(t *testing.T) {
   {
      b, err := BandID("2853020814")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", b)
   }
   {
      b, err := BandURL("duststoredigital.com")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", b)
   }
}
