package bandcamp

import (
   "fmt"
   "testing"
)

const bandID = 2853020814

func TestBand(t *testing.T) {
   Verbose(true)
   ban := new(Band)
   if err := ban.Post(bandID); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", ban)
}
