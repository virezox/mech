package vimeo

import (
   "fmt"
   "testing"
)

func TestVideo(t *testing.T) {
   vid, err := NewVideo(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
