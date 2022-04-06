package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestPlayer(t *testing.T) {
   for name, version := range names {
      play, err := newPlayer(name, version)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(play.PlayabilityStatus.Status, name)
      time.Sleep(time.Second)
   }
}
