package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "testing"
)

func TestAndroid(t *testing.T) {
   a, err := youtube.NewAndroid("XeojXq6ySs4")
   if err != nil {
      t.Fatal(err)
   }
   a.StreamingData.AdaptiveFormats.Sort()
   for _, f := range a.StreamingData.AdaptiveFormats {
      fmt.Println(f.Height, f.MimeType, f.Bitrate)
   }
}
