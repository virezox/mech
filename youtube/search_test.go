package youtube

import (
   "fmt"
   "testing"
)

func TestSearch(t *testing.T) {
   Verbose(true)
   s, err := NewSearch("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, i := range s.Items() {
      fmt.Println(i.CompactVideoRenderer)
   }
}
