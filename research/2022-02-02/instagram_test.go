package instagram

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   // image:1 video:0
   "CZVEugIPkVn",
   // image:6 video:0
   "CXzGW6RPNmy",
   // image:0 video:3 DASH:0
   "BQ0eAlwhDrw",
   // image:2 video:1 DASH:1
   "CUK-1wjqqsP",
   // image:0 video:1 DASH:1
   "CLHoAQpCI2i",
}

func TestMedia(t *testing.T) {
   for i, test := range tests {
      if i >= 1 {
         fmt.Println("---")
      }
      items, err := MediaItems(test)
      if err != nil {
         t.Fatal(err)
      }
      for _, item := range items {
         for i, info := range item.Infos() {
            if i >= 1 {
               fmt.Println("---")
            }
            addr, err := info.URL()
            if err != nil {
               t.Fatal(err)
            }
            fmt.Printf("%q\n", addr)
         }
      }
      time.Sleep(time.Second)
   }
}
