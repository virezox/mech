package vimeo

import (
   "fmt"
   "testing"
)

func TestVimeo(t *testing.T) {
   v, err := newVideos()
   if err != nil {
      t.Fatal(err)
   }
   p, err := v.playground("66531465")
   if err != nil {
      t.Fatal(err)
   }
   c, err := v.callable(p)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", c)
}
