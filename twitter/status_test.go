package twitter

import (
   "fmt"
   "testing"
)

const statusID = 1470124083547418624

var guest = &Guest{"1475108770955022337"}

func TestStatus(t *testing.T) {
   stat, err := NewStatus(guest, statusID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stat)
}
