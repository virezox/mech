package vimeo

import (
   "fmt"
   "testing"
)

var clip = Clip{581039021, 9603038895}

var videos = []string{
   "https://vimeo.com/66531465",
   "https://vimeo.com/477957994/2282452868",
   "https://vimeo.com/477957994?unlisted_hash=2282452868",
}

func TestVimeo(t *testing.T) {
   web, err := NewJsonWeb()
   if err != nil {
      t.Fatal(err)
   }
   LogLevel = 1
   video, err := web.Video(&clip)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(video)
}

func TestScan(t *testing.T) {
   for _, video := range videos {
      clip, err := NewClip(video)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", clip)
   }
}
