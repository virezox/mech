package vimeo

import (
   "fmt"
   "testing"
)

const path = "/videos/581039021:9603038895"

func TestVimeo(t *testing.T) {
   logLevel = 1
   web, err := NewJsonWeb()
   if err != nil {
      t.Fatal(err)
   }
   video, err := web.Video(path)
   if err != nil {
      t.Fatal(err)
   }
   for _, down := range video.Download {
      fmt.Printf("%+v\n", down)
   }
}

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
