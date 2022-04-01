package abc

import (
   "os"
   "testing"
)

const grey =
   "/shows/greys-anatomy/episode-guide/season-18/12-the-makings-of-you"

func TestMech(t *testing.T) {
   res, err := Route(grey)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
