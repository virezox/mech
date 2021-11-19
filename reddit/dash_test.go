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
   d, err := l.DASH()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", d)
}
