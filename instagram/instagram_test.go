package instagram

import (
   "fmt"
   "testing"
   "time"
)

const id = "CT-cnxGhvvO"

func TestInsta(t *testing.T) {
   Verbose = true
   for range [16]struct{}{} {
      c, err := NewSidecar(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(c)
      time.Sleep(time.Second)
   }
}
