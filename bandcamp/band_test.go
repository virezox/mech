package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

const bandID = 2853020814

func TestBand(t *testing.T) {
   Verbose(true)
   ban := new(Band)
   if err := ban.Get(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ban)
   time.Sleep(100 * time.Millisecond)
   if err := ban.Post(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ban)
   time.Sleep(100 * time.Millisecond)
   if err := ban.PostForm(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ban)
}
