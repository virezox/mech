package youtube_test

import (
   "github.com/89z/mech/youtube"
   "testing"
)

func TestFormat(t *testing.T) {
   a, err := youtube.NewAndroid("eAzIAjTBGgU")
   if err != nil {
      t.Fatal(err)
   }
   // this should fail
   f, err := a.NewFormat(247)
   if err == nil {
      t.Fatalf("%+v\n", f)
   }
}
