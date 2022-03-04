package paramount

import (
   "fmt"
   "testing"
)

func TestParamount(t *testing.T) {
   seg, err := segment()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", seg.Key)
}
