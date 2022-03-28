package pbs

import (
   "fmt"
   "testing"
)

const frontlineTest =
   "https://www.pbs.org/wgbh/frontline/film/inside-italys-covid-war/"

func TestFrontline(t *testing.T) {
   line, err := NewFrontline(frontlineTest)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", line)
}
