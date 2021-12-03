package bbc

import (
   "fmt"
   "testing"
)

const id = 10462520

func TestNews(t *testing.T) {
   item, err := NewNewsVideo(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
