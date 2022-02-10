package vimeo

import (
   "fmt"
   "testing"
)

var videos = []string{
   "https://vimeo.com/66531465",
   "https://vimeo.com/477957994/2282452868",
   "https://vimeo.com/477957994?unlisted_hash=2282452868",
}

func TestScan(t *testing.T) {
   for _, video := range videos {
      clip, err := newClip(video)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", clip)
   }
}
