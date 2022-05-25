package bbcamerica

import (
   "fmt"
   "testing"
)

func TestUnauth(t *testing.T) {
   auth, err := NewUnauth()
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
