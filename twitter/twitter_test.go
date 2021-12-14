package twitter

import (
   "fmt"
   "testing"
)

const id = 1470124083547418624

func TestTwitter(t *testing.T) {
   act, err := newActivate()
   if err != nil {
      t.Fatal(err)
   }
   stat, err := act.status(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stat)
}
