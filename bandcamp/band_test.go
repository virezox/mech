package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const bandID = "2853020814"

func TestBand(t *testing.T) {
   Verbose(true)
   b := new(Band)
   if err := b.Get(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   time.Sleep(100 * time.Millisecond)
   if err := b.Post(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
   time.Sleep(100 * time.Millisecond)
   if err := b.PostForm(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", b)
}
