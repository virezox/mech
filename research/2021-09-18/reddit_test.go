package reddit

import (
   "fmt"
   "testing"
)

const id = "pqfqoy"

func TestReddit(t *testing.T) {
   Verbose = true
   lists, err := listings(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", lists)
}
