package pbs

import (
   "fmt"
   "testing"
)

const frontline =
   "https://www.pbs.org/wgbh/frontline/film/inside-italys-covid-war/"

func TestFrontline(t *testing.T) {
   front, err := NewWidgeter(frontline)
   if err != nil {
      t.Fatal(err)
   }
   widget, err := front.Widget()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", widget)
}
