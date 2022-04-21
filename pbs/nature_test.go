package pbs

import (
   "fmt"
   "testing"
)

const nature = "https://www.pbs.org/wnet/nature/about-american-horses/"

func TestNature(t *testing.T) {
   nat, err := NewWidgeter(nature)
   if err != nil {
      t.Fatal(err)
   }
   widget, err := nat.Widget()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", widget)
}
