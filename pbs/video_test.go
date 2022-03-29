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
      video, err := NewWidgeter(test)
      if err != nil {
         t.Fatal(err)
      }
      widget, err := video.Widget()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", widget)
      time.Sleep(time.Second)
   }
}
