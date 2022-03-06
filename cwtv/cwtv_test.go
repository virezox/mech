package cwtv

import (
   "fmt"
   "testing"
   "time"
)

var shows = []string{
   "https://cwtv.com/shows/4400/group-efforts/?play=deec61a8-e0a1-4c01-8906-4e0b363350d5",
   "https://cwtv.com/shows/4400/past-is-prologue/?play=dead9843-33b1-4201-adf8-310692fe147f",
   "https://cwtv.com/shows/4400/present-is-prologue/?play=6cc4708a-9b9e-45e2-ada4-468355f6cb38",
}

func TestMedia(t *testing.T) {
   for _, show := range shows {
      play, err := GetPlay(show)
      if err != nil {
         t.Fatal(err)
      }
      addr, err := Media(play)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(addr)
      time.Sleep(time.Second)
   }
}
