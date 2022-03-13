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
   check, err := Clip{ID: id}.Check(password)
   if err != nil {
      t.Fatal(err)
   }
   for _, pro := range check.Request.Files.Progressive {
      fmt.Printf("%+v\n", pro.WithURL(""))
   }
}
