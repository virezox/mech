package youtube

import (
   "fmt"
   "testing"
)

func TestContinue(t *testing.T) {
   s, err := encode(newContinuation("q5UnT4Ik6KU"))
   if err != nil {
      t.Fatal(err)
   }
   if s != "Eg0SC3E1VW5UNElrNktVGAYyDyINIgtxNVVuVDRJazZLVQ==" {
      t.Fatal(s)
   }
}

func TestParam(t *testing.T) {
   var p param
   p.video()
   s, err := encode(p)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(s)
   /*
   m := youtube.Params["TYPE"]["Video"]
   if s := m.Encode(); s != "EgIQAQ==" {
   */
}
