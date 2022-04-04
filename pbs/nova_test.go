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
      nova, err := NewWidgeter(test)
      if err != nil {
         t.Fatal(err)
      }
      widget, err := nova.Widget()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%a\n", widget)
      time.Sleep(time.Second)
   }
}
