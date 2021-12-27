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
   dash, err := con.DASH()
   if err != nil {
      t.Fatal(err)
   }
   for _, vid := range dash.Video {
      fmt.Printf("%+v\n", vid)
   }
}
