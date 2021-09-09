package googleplay

import (
   "fmt"
   "testing"
)

var ids = []uint16{0x1302, 0x1303}

func TestCipher(t *testing.T) {
   c, err := newCipherSuites()
   if err != nil {
      t.Fatal(err)
   }
   for _, id := range ids {
      fmt.Printf("0x%04X, // %v\n", id, c.get(id))
   }
}
