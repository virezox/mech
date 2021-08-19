package js

import (
   "fmt"
   "os"
   "testing"
)

func TestJS(t *testing.T) {
   b, err := os.ReadFile("index.js")
   if err != nil {
      t.Fatal(err)
   }
   for k, v := range newValues(b) {
      fmt.Printf("%v\n%s\n", k, v)
   }
}
