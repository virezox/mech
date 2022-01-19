package vimeo

import (
   "fmt"
   "testing"
)

func TestConfig(t *testing.T) {
   con, err := NewConfig(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", con)
}
