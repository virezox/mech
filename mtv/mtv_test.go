package mtv

import (
   "fmt"
   "testing"
   "time"
)

type propertyTest struct {
   typ, id string
}

var tests = []propertyTest{
   {"episode", "scyb0g"},
   {"showvideo", "s5iqyc"},
}

func TestProperty(t *testing.T) {
   for _, test := range tests {
      prop, err := NewProperty(test.typ, test.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", prop)
      time.Sleep(time.Second)
   }
}

func TestTopaz(t *testing.T) {
   prop, err := NewProperty(tests[0].typ, tests[0].id)
   if err != nil {
      t.Fatal(err)
   }
   top, err := prop.Topaz()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", top)
}
