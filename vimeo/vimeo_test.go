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
   vids, err := con.Videos()
   if err != nil {
      t.Fatal(err)
   }
   for _, vid := range vids {
      fmt.Println(vid.URL())
   }
}

func TestEmbed(t *testing.T) {
   emb, err := NewEmbed(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", emb)
}
