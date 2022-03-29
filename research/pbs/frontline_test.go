package pbs

import (
   "fmt"
   "testing"
   "time"
)

var frontlineTests = []string{
   "https://www.pbs.org/video/frontlineworld-children-of-the-taliban/",
   "https://www.pbs.org/wgbh/frontline/film/inside-italys-covid-war/",
   "https://www.pbs.org/wgbh/masterpiece/episodes/downton-abbey-s6-e2/",
   "https://www.pbs.org/wnet/nature/about-american-horses/",
}

func TestFrontline(t *testing.T) {
   for _, test := range frontlineTests {
      line, err := NewFrontline(test)
      if err != nil {
         t.Fatal(err)
      }
      bridge, err := line.VideoObject().VideoBridge()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(bridge)
      time.Sleep(time.Second)
   }
}
