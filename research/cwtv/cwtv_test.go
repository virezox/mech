package cwtv

import (
   "os"
   "testing"
)

const play = "deec61a8-e0a1-4c01-8906-4e0b363350d5"

func TestMedia(t *testing.T) {
   LogLevel = 1
   res, err := media(play)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
