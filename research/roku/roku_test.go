package roku

import (
   "fmt"
   "testing"
)

func TestRoku(t *testing.T) {
   site, err := newCrossSite()
   if err != nil {
      t.Fatal(err)
   }
   play, err := site.playback()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
