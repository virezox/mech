package pbs

import (
   "fmt"
   "testing"
)

const videoTest = "https://www.pbs.org/wnet/nature/about-american-horses/"

func TestVideo(t *testing.T) {
   video, err := NewVideo(videoTest)
   if err != nil {
      t.Fatal(err)
   }
   bridge, err := video.Bridge()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", bridge)
}
