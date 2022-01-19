package vimeo

import (
   "fmt"
   "testing"
)

const id = 660408476

func TestConfig(t *testing.T) {
   con, err := NewConfig(id)
   if err != nil {
      t.Fatal(err)
   }
   mas, err := con.Master()
   if err != nil {
      t.Fatal(err)
   }
   for _, vid := range mas.Video {
      fmt.Println(vid.URL(con))
   }
}

func TestOembed(t *testing.T) {
   emb, err := NewOembed(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", emb)
}
