package instagram

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   // type:1 video:0 image:1
   "CZVEugIPkVn",
   // type:2 video:1 image:0
   "CUDJ4YhpF0Z",
   // type:8 video:0 image:6
   "CXzGW6RPNmy",
   // type:8 video:1 image:2
   "CUK-1wjqqsP",
   // type:8 video:3 image:0
   "BQ0eAlwhDrw",
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
            fmt.Printf("%q\n", info.URL())
         }
      }
      time.Sleep(time.Second)
   }
}
