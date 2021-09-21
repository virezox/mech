package instagram

import (
   "fmt"
   "testing"
)

func TestInsta(t *testing.T) {
   Verbose = true
   c, err := NewSidecar("CT-cnxGhvvO")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", c)
}
