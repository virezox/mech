package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestSort(t *testing.T) {
   p, err := youtube.PlayerAndroid("XeojXq6ySs4")
   if err != nil {
      t.Fatal(err)
   }
   p.AdaptiveFormats.Sort()
   for _, f := range p.AdaptiveFormats {
      fmt.Println(f.Height, f.MimeType, f.Bitrate)
   }
}
