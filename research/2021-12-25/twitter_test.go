package twitter

import (
   "fmt"
   "testing"
)

const id = 1468949573565657094

func TestTwitter(t *testing.T) {
   act, err := NewActivate()
   if err != nil {
      t.Fatal(err)
   }
   stat, err := act.Status(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stat)
}
