package pbs

import (
   "fmt"
   "testing"
)

const masterpiece =
   "https://www.pbs.org/wgbh/masterpiece/episodes/downton-abbey-s6-e2/"

func TestMasterpiece(t *testing.T) {
   master, err := NewMasterpiece(masterpiece)
   if err != nil {
      t.Fatal(err)
   }
   wid, err := master.Widget()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(wid)
}
