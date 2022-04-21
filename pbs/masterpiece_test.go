package pbs

import (
   "fmt"
   "testing"
)

const masterpiece =
   "https://www.pbs.org/wgbh/masterpiece/episodes/downton-abbey-s6-e2/"

func TestMasterpiece(t *testing.T) {
   master, err := NewWidgeter(masterpiece)
   if err != nil {
      t.Fatal(err)
   }
   widget, err := master.Widget()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%a\n", widget)
}
