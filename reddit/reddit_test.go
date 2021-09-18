package reddit

import (
   "fmt"
   "testing"
)

const id = "pqfqoy"

func TestReddit(t *testing.T) {
   Verbose = true
   p, err := NewPost(id)
   if err != nil {
      t.Fatal(err)
   }
   t3, err := p.T3()
   if err != nil {
      t.Fatal(err)
   }
   m, err := t3.MPD()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
