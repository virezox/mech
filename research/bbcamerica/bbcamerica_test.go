package bbcamerica

import (
   "os"
   "testing"
)

func TestUnauth(t *testing.T) {
   auth, err := NewUnauth()
   if err != nil {
      t.Fatal(err)
   }
   res, err := auth.Playback()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
