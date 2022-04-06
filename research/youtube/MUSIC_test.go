package youtube

import (
   "testing"
)

func TestMusicIntegration(t *testing.T) {
   const name = "MUSIC_INTEGRATIONS"
   err := post(name, names[name])
   if err != nil {
      t.Fatal(err)
   }
}
