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

func TestNova(t *testing.T) {
   for _, test := range playerTests {
      nova, err := NewNova(test)
      if err != nil {
         t.Fatal(err)
      }
      wid, err := nova.Episode().Asset().Widget()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", wid)
      time.Sleep(time.Second)
   }
}
