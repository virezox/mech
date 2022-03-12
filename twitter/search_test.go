package twitter

import (
   "fmt"
   "testing"
)

func TestSearch(t *testing.T) {
   guest, err := NewGuest()
   if err != nil {
      t.Fatal(err)
   }
   LogLevel = 1
   sea, err := guest.Search("filter:spaces", 2)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", sea)
}
