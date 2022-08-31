package vimeo

import (
   "fmt"
   "testing"
)

func Test_Web(t *testing.T) {
   web, err := New_JSON_Web()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", web)
}
