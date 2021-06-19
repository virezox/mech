package youtube_test

import (
   "github.com/89z/mech/youtube"
   "io"
   "testing"
)

const desc = "Provided to YouTube by Epitaph\n\nSnowflake · Kate Bush\n\n" +
"50 Words For Snow\n\n" +
"℗ Noble & Brite Ltd. trading as Fish People, under exclusive license to Anti Inc.\n\n" +
"Released on: 2011-11-22\n\nMusic  Publisher: Noble and Brite Ltd.\n" +
"Composer  Lyricist: Kate Bush\n\nAuto-generated by YouTube."

func TestSignatureCipher(t *testing.T) {
   v, err := youtube.NewPlayer("XeojXq6ySs4")
   if err != nil {
      t.Fatal(err)
   }
   if v.Description() != desc {
      t.Fatalf("%+v\n", v)
   }
   if v.ViewCount() == 0 {
      t.Fatalf("%+v\n", v)
   }
   f, err := v.NewFormat(249)
   if err != nil {
      t.Fatal(err)
   }
   b, err := youtube.NewBaseJS()
   if err != nil {
      t.Fatal(err)
   }
   if err := b.Get(); err != nil {
      t.Fatal(err)
   }
   if err := f.Write(io.Discard); err != nil {
      t.Fatal(err)
   }
}

func TestContentLength(t *testing.T) {
   v, err := youtube.NewPlayer("eAzIAjTBGgU")
   if err != nil {
      t.Fatal(err)
   }
   // this should fail
   f, err := v.NewFormat(247)
   if err == nil {
      t.Fatal(f)
   }
}
