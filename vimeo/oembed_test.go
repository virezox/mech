package vimeo

import (
   "fmt"
   "testing"
)

const id = 660408476

func TestEmbed(t *testing.T) {
   emb, err := NewEmbed(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", emb)
}
