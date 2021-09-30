package instagram

import (
   "fmt"
   "testing"
)

const code = "CT-cnxGhvvO"

func TestSidecar(t *testing.T) {
   m, err := NewMedia(code)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
