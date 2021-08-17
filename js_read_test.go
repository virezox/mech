package mech

import (
   "fmt"
   "strings"
   "testing"
)

const d = `d={ab:9,'cd':9,'c"d':9,"ef":9,"e'f":9}`

func TestJsRead(t *testing.T) {
   r, err := NewJsReader(strings.NewReader(d))
   if err != nil {
      t.Fatal(err)
   }
   be := make(map[string]string)
   r.Read(be)
   for k, v := range be {
      fmt.Printf("%q\n%v\n", k, v)
   }
}
