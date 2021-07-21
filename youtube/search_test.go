package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestSearch(t *testing.T) {
   s, err := youtube.NewSearch("nelly furtado say it right")
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range s.Videos() {
      fmt.Printf("%+v\n", v)
   }
}
