package js

import (
   "encoding/json"
   "fmt"
   "testing"
)

func TestValues(t *testing.T) {
   b := []byte(`d={ab:9,'cd':9,'c"d':9,"ef":9,"e'f":9}`)
   v, err := Parse(b)
   if err != nil {
      t.Fatal(err)
   }
   var m map[string]int
   if err := json.Unmarshal(v.Get("d"), &m); err != nil {
      t.Fatal(err)
   }
   fmt.Println(m)
}
