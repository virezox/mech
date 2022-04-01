package abc

import (
   "fmt"
   "testing"
)

const grey =
   "/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you"

func TestMech(t *testing.T) {
   route, err := NewRoute(grey)
   if err != nil {
      t.Fatal(err)
   }
   play, err := route.Player()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
