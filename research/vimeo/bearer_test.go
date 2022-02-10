package vimeo

import (
   "fmt"
   "testing"
)

func TestBearer(t *testing.T) {
   bear, err := NewBearer()
   if err != nil {
      t.Fatal(err)
   }
   video, err := bear.Video(path)
   if err != nil {
      t.Fatal(err)
   }
   for _, down := range video.Download {
      fmt.Printf("%+v\n", down)
   }
}
