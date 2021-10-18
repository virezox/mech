package bandcamp

import (
   "fmt"
   "testing"
)

const label = "https://multiculti.bandcamp.com"

func TestBand(t *testing.T) {
   Verbose(true)
   i, err := NewItem(label)
   if err != nil {
      t.Fatal(err)
   }
   l, err := NewBand(i.Item_ID)
   if err != nil {
      t.Fatal(err)
   }
   b, err := NewBand(l.Artists[0].ID)
   if err != nil {
      t.Fatal(err)
   }
   a, err := b.Discography[0].Tralbum()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", a)
}
