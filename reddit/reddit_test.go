package reddit

import (
   "fmt"
   "testing"
)

const id = "pqfqoy"

func TestReddit(t *testing.T) {
   p, err := NewPost(id)
   if err != nil {
      t.Fatal(err)
   }
   l, err := p.Link()
   if err != nil {
      t.Fatal(err)
   }
   m, err := l.MPD()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}
