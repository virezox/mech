package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestSort(t *testing.T) {
   p, err := youtube.NewPlayer("XeojXq6ySs4", youtube.Key, youtube.Android)
   if err != nil {
      t.Fatal(err)
   }
   p.StreamingData.AdaptiveFormats.Sort()
   for _, f := range p.StreamingData.AdaptiveFormats {
      fmt.Println(f.Height, f.MimeType, f.Bitrate)
   }
}
