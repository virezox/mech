package vimeo

import (
   "fmt"
   "testing"
)

func TestBearer(t *testing.T) {
   logLevel = 1
   bear, err := newBearer()
   if err != nil {
      t.Fatal(err)
   }
   video, err := bear.video(path)
   if err != nil {
      t.Fatal(err)
   }
   for _, down := range video.Download {
      fmt.Printf("%+v\n", down)
   }
}
