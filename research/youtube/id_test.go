package youtube

import (
   "fmt"
   "testing"
)

var addrs = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

func TestID(t *testing.T) {
   for _, addr := range addrs {
      id, err := VideoID(addr)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", id)
   }
}
