package tomato

import (
   "fmt"
   "testing"
)

func TestMovie(t *testing.T) {
   m, err := NewMovie("https://www.rottentomatoes.com/m/one_night_in_miami")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
