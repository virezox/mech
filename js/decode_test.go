package js

import (
   "fmt"
   "strings"
   "testing"
)

const s = `d={ab:9,'cd':9,'c"d':9,"ef":9,"e'f":9}`

func TestDecode(t *testing.T) {
   d, err := NewDecoder(strings.NewReader(s))
   if err != nil {
      t.Fatal(err)
   }
   be := make(map[string]string)
   d.Decode(be)
   for k, v := range be {
      fmt.Printf("%q\n%v\n", k, v)
   }
}
