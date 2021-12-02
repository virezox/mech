package bbc

import (
   "fmt"
   "testing"
)

const (
   externalID = "p0b7p8sq"
   id = 10462520
)

func TestSelector(t *testing.T) {
   sel, err := NewSelector(externalID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", sel)
}

func TestNews(t *testing.T) {
   item, err := NewNewsVideo(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
