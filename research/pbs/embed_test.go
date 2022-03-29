package pbs

import (
   "fmt"
   "testing"
   "time"
)

var embedTests = []string{
   "https://www.pbs.org/wgbh/frontline/film/inside-italys-covid-war/",
   "https://www.pbs.org/wgbh/masterpiece/episodes/downton-abbey-s6-e2/",
}

func TestEmbed(t *testing.T) {
   for _, test := range embedTests {
      line, err := NewEmbed(test)
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
