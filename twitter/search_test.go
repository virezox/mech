package twitter

import (
   "fmt"
   "testing"
)

func TestSearch(t *testing.T) {
   sea, err := NewSearch("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", sea)
}
