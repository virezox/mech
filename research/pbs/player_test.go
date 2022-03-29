package pbs

import (
   "fmt"
   "testing"
   "time"
)

var playerTests = []string{
   "https://www.pbs.org/wgbh/nova/video/australias-first-4-billion-years-awakening/",
   "https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/",
   "https://www.pbs.org/wgbh/nova/video/the-planets-inner-worlds/",
}

func TestPlayer(t *testing.T) {
   for _, test := range playerTests {
      play, err := NewPlayer(test)
      if err != nil {
         t.Fatal(err)
      }
      bridge, err := play.Episode().Asset().Bridge()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", bridge)
      time.Sleep(time.Second)
   }
}
