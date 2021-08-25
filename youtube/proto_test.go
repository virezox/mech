package youtube_test

import (
   "github.com/89z/mech/youtube"
   "testing"
)

func TestContinue(t *testing.T) {
   s := youtube.Continuation("q5UnT4Ik6KU").Encode()
   if s != "Eg0SC3E1VW5UNElrNktVGAYyDyINIgtxNVVuVDRJazZLVQ==" {
      t.Fatal(s)
   }
}

func TestParam(t *testing.T) {
   m := youtube.Params["TYPE"]["Video"]
   if s := m.Encode(); s != "EgIQAQ==" {
      t.Fatal(s)
   }
}
