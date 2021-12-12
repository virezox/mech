package vimeo

import (
   "fmt"
   "testing"
)

const id = 66531465

func TestConfig(t *testing.T) {
   con, err := NewConfig(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", con)
}

func TestVideo(t *testing.T) {
   vid, err := NewVideo(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
