package js

import (
   "fmt"
   "testing"
)

var b = []byte(`
a={"bc":9,"de":9};
f={"gh":9,"ij":9};
`)

func TestJS(t *testing.T) {
   for k, v := range newValues(b) {
      fmt.Printf("%v %s\n", k, v)
   }
}
