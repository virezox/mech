package mtv

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestTopaz(t *testing.T) {
   res, err := topaz()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

type propertyTest struct {
   typ, id string
}

var tests = []propertyTest{
   {"episode", "scyb0g"},
   {"showvideo", "s5iqyc"},
}

func TestProperty(t *testing.T) {
   for _, test := range tests {
      prop, err := newProperty(test.typ, test.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", prop)
      time.Sleep(time.Second)
   }
}
