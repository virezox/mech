package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestSort(t *testing.T) {
   a, err := youtube.NewAndroid("XeojXq6ySs4")
   if err != nil {
      t.Fatal(err)
   }
   a.AdaptiveFormats.Sort()
   for _, f := range a.AdaptiveFormats {
      fmt.Println(f.Height, f.MimeType, f.Bitrate)
   }
}
