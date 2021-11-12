package youtube

import (
   "fmt"
   "testing"
)

func TestSearch(t *testing.T) {
   s, err := NewSearch("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, i := range s.Items() {
      fmt.Println(i.CompactVideoRenderer)
   }
}
