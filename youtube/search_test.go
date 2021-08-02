package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestSearch(t *testing.T) {
   youtube.Verbose = true
   s, err := youtube.NewSearch("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   for _, i := range s.Items() {
      fmt.Println(i.CompactVideoRenderer)
   }
}
