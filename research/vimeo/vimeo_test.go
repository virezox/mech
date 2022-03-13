package vimeo

import (
   "fmt"
   "testing"
)

const (
   id = 412573977
   password = "butter"
)

func TestVideo(t *testing.T) {
   vid, err := NewVideo(id, password)
   if err != nil {
      t.Fatal(err)
   }
   for _, pro := range vid.Request.Files.Progressive {
      fmt.Printf("%+v\n", pro)
   }
}
