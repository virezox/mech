package pbs

import (
   "fmt"
   "testing"
   "time"
)

var contents = []string{
   "https://www.pbs.org/video/frontlineworld-children-of-the-taliban/",
   "https://www.pbs.org/video/the-future-of-restaurants-rbs0v8/",
}

func TestContent(t *testing.T) {
   for _, content := range contents {
      con, err := NewContent(content)
      if err != nil {
         t.Fatal(err)
      }
      bridge, err := con.Bridge()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(bridge)
      time.Sleep(time.Second)
   }
}
