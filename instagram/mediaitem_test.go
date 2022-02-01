package instagram

import (
   "fmt"
   "testing"
   "time"
)

type testType struct {
   shortcode string
   id uint64
}

var tests = []testType{
   // type:1 video:0 image:1
   {"CZVEugIPkVn", 2762134734241678695},
   // type:2 video:1 image:0
   {"CUDJ4YhpF0Z", 2667018861376986393},
   // type:8 video:0 image:6
   {"CXzGW6RPNmy", 2734557361417804210},
   // type:8 video:1 image:2
   {"CUK-1wjqqsP", 2669222102324390671},
   // type:8 video:3 image:0
   {"BQ0eAlwhDrw", 1455920561485265648},
}

func TestMedia(t *testing.T) {
   for _, test := range tests {
      items, err := MediaItems(test.id)
      if err != nil {
         t.Fatal(err)
      }
      for _, item := range items {
         for _, info := range item.Infos() {
            ver, err := info.Version()
            if err != nil {
               t.Fatal(err)
            }
            fmt.Println(ver.URL[:99])
         }
      }
      time.Sleep(time.Second)
   }
}

func TestID(t *testing.T) {
   for _, test := range tests {
      id, err := GetID(test.shortcode)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}

