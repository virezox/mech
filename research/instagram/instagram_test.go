package instagram

import (
   "fmt"
   "testing"
)

const shortcode = "CZVEugIPkVn"

func TestMedia(t *testing.T) {
   items, err := MediaItems(shortcode)
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range items {
      fmt.Printf("%+v\n", item)
   }
}
