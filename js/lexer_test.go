package js

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func TestLexer(t *testing.T) {
   b, err := os.ReadFile("index.js")
   if err != nil {
      t.Fatal(err)
   }
   v := NewLexer(b).Values()
   var a []interface{}
   if err := json.Unmarshal(v["apps"], &a); err != nil {
      t.Fatal(err)
   }
   fmt.Println(a)
}
