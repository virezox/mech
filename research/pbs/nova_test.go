package pbs

import (
   "fmt"
   "testing"
   "time"
)

var novaTests = []string{
   "https://www.pbs.org/wgbh/nova/video/australias-first-4-billion-years-awakening/",
   "https://www.pbs.org/wgbh/nova/video/nova-universe-revealed-milky-way/",
   "https://www.pbs.org/wgbh/nova/video/the-planets-inner-worlds/",
}

func TestNova(t *testing.T) {
   for _, test := range novaTests {
      data, err := NewNextData(test)
      if err != nil {
         t.Fatal(err)
      }
      video, err := data.Episode().Asset().VideoBridge()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
