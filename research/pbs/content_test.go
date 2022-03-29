package pbs

import (
   "fmt"
   "testing"
   "time"
)

const contentTest =
   "https://www.pbs.org/video/frontlineworld-children-of-the-taliban/"

func TestFrontline(t *testing.T) {
   line, err := NewFrontline(test)
   if err != nil {
      t.Fatal(err)
   }
   bridge, err := line.VideoObject().VideoBridge()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(bridge)
}
