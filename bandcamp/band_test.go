package bandcamp

import (
   "fmt"
   "testing"
)

func TestBand(t *testing.T) {
   {
      b, err := BandGet("2853020814")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", b)
   }
   {
      b, err := BandPost("2853020814")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", b)
   }
}
