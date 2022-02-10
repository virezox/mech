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
