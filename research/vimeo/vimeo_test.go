package vimeo

import (
   "fmt"
   "testing"
)

// vimeo.com/_rv/title?path=/581039021/9603038895
var videos = []string{
   // vimeo.com/581039021/9603038895
   "/videos/581039021:9603038895",
   // vimeo.com/660408476
   "/videos/660408476",
}

func TestVideo(t *testing.T) {
   web, err := newJsonWeb()
   if err != nil {
      t.Fatal(err)
   }
   logLevel = 1
   video, err := web.video(videos[0])
   if err != nil {
      t.Fatal(err)
   }
   for _, down := range video.Download {
      fmt.Printf("%+v\n", down)
   }
}
