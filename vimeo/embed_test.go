package vimeo

import (
   "fmt"
   "testing"
)

func TestVideo(t *testing.T) {
   emb, err := NewEmbed(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", emb)
}
