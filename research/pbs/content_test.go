package pbs

import (
   "fmt"
   "testing"
)

const contentTest =
   "https://www.pbs.org/video/frontlineworld-children-of-the-taliban/"

func TestContent(t *testing.T) {
   con, err := NewContent(contentTest)
   if err != nil {
      t.Fatal(err)
   }
   brid, err := con.Bridge()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(brid)
}
