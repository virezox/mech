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
   fmt.Printf("%+v\n", p)
   media, err := p.MPD()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
