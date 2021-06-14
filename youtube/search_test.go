package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
   "time"
)

func TestSearch(t *testing.T) {
   for _, query := range []string{"radiohead", "nelly furtado afraid"} {
      s, err := youtube.NewSearch(query)
      if err != nil {
         t.Fatal(err)
      }
      for _, v := range s.VideoRenderers() {
         fmt.Printf("%+v\n", v)
      }
      time.Sleep(time.Second)
   }
}
