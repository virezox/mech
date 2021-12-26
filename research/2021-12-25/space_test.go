package twitter

import (
   "fmt"
   "testing"
)

const spaceID = "1OdKrBnaEPXKX"

func TestSpace(t *testing.T) {
   space, err := guest.Space(spaceID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", space)
}
