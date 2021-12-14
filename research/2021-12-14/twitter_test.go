package twitter

import (
   "fmt"
   "testing"
)

func TestTwitter(t *testing.T) {
   act, err := newActivate()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", act)
}
