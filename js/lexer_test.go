package js

import (
   "fmt"
   "os"
   "testing"
)

func TestLexer(t *testing.T) {
   b, err := os.ReadFile("index.js")
   if err != nil {
      t.Fatal(err)
   }
   for k, v := range NewLexer(b).Values() {
      fmt.Printf("%q\n%s\n\n", k, v)
   }
}
