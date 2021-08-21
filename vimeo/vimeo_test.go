package vimeo

import (
   "fmt"
   "testing"
)

func TestConfig(t *testing.T) {
   c, err := NewConfig("66531465")
   if err != nil {
      t.Fatal(err)
   }
   v, err := c.Video()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", v)
}
