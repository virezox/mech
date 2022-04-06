package youtube

import (
   "fmt"
   "testing"
)

func TestMusicIntegration(t *testing.T) {
   const (
      name = "MUSIC_INTEGRATIONS"
      version = "0.1"
   )
   res, err := post(name, version)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Println(res.Status, name, version)
}
