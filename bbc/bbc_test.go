package bbc

import (
   "fmt"
   "testing"
)

const addr = "https://www.bbc.com/news/av/10462520"

func TestNews(t *testing.T) {
   item, err := NewNewsVideo(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
