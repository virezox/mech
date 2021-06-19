package youtube

import (
   "fmt"
   "testing"
)

func TestPlayer(t *testing.T) {
   play, err := NewPlayer("9cNrM5AIigw")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
