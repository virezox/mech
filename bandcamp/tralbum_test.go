package bandcamp

import (
   "fmt"
   "testing"
)

var details = []Detail{
   {1, 79940049, "a"},
   {1, 2809477874, "t"},
}

func TestTralbum(t *testing.T) {
   Verbose(true)
   for _, detail := range details {
      tra, err := detail.Tralbum()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", tra)
   }
}
