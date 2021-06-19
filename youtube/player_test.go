package youtube_test

import (
   "github.com/89z/mech/youtube"
   "testing"
)

func TestPlayer(t *testing.T) {
   err := youtube.NewPlayer("NMYIVsdGfoo")
   if err != nil {
      t.Fatal(err)
   }
}
