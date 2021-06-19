package youtube_test

import (
   "github.com/89z/mech/youtube"
   "io"
   "testing"
)

func TestSignatureCipher(t *testing.T) {
   {
      v, err := youtube.NewPlayer("XeojXq6ySs4")
      if err != nil {
         t.Fatal(err)
      }
      f, err := v.NewFormat(249)
      if err != nil {
         t.Fatal(err)
      }
      if err := f.Write(io.Discard); err != nil {
         t.Fatal(err)
      }
   }
   {
      v, err := youtube.YouTubeI("XeojXq6ySs4")
      if err != nil {
         t.Fatal(err)
      }
      f, err := v.NewFormat(249)
      if err != nil {
         t.Fatal(err)
      }
      if err := f.Write(io.Discard); err != nil {
         t.Fatal(err)
      }
   }
}
