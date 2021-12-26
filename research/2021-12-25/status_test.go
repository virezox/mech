package twitter

import (
   "fmt"
   "testing"
)

const statusID = 1470124083547418624

var guest = Guest{"1474966304180297735"}

func TestStatus(t *testing.T) {
   stat, err := guest.Status(statusID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stat)
}
