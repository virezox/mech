package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestSearch(t *testing.T) {
   r, err := youtube.NewSearch("nelly furtado say it right").Post()
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range r.VideoRenderers() {
      fmt.Printf("%+v\n", v)
   }
}
