package vimeo

import (
   "fmt"
   "testing"
)

const id = "66531465"

func TestConfig(t *testing.T) {
   c, err := NewConfig(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", c)
}

func TestVideo(t *testing.T) {
   v, err := NewVideo(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", v)
}
