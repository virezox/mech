package youtube_test

import (
   "github.com/89z/mech/youtube"
   "testing"
)

func TestContinue(t *testing.T) {
   s, err := youtube.NewContinuation("q5UnT4Ik6KU").Encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "Eg0SC3E1VW5UNElrNktVGAYyDyINIgtxNVVuVDRJazZLVQ==" {
      t.Fatal(s)
   }
}

func TestParam(t *testing.T) {
   var p youtube.Param
   p.Video()
   s, err := p.Encode()
   if err != nil {
      t.Fatal(err)
   }
   if s != "EgIQAQ==" {
      t.Fatal(s)
   }
}
