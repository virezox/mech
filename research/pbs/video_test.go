package pbs

import (
   "fmt"
   "testing"
   "time"
)

var videoTests = []string{
   "https://www.pbs.org/video/frontlineworld-children-of-the-taliban/",
   "https://www.pbs.org/video/the-future-of-restaurants-rbs0v8/",
}

func TestVideo(t *testing.T) {
   for _, test := range videoTests {
      vid, err := NewVideo(test)
      if err != nil {
         t.Fatal(err)
      }
      wid, err := vid.Widget()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(wid)
      time.Sleep(time.Second)
   }
}
