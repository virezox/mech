package pbs

import (
   "fmt"
   "testing"
)

const frontline =
   "https://www.pbs.org/wgbh/frontline/film/inside-italys-covid-war/"

func TestFrontline(t *testing.T) {
   front, err := NewFrontline(frontline)
   if err != nil {
      t.Fatal(err)
   }
   wid, err := front.Widget()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(wid)
}
