package bandcamp

import (
   "fmt"
   "testing"
)

func TestBand(t *testing.T) {
   b := new(Band)
   if err := b.Get(2853020814); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   if err := b.Post(2853020814); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   bi := new(BandInfo)
   if err := bi.Get(2853020814); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", bi)
   if err := bi.Post(2853020814); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", bi)
}
