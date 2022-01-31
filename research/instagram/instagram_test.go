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
   // 0 image 1 video
   {"CUDJ4YhpF0Z", 2667018861376986393},
   // 0 image 2 video
   {"BQ0eAlwhDrw", 1455920561485265648},
   // 1 image 0 video
   {"CZVEugIPkVn", 2762134734241678695},
   // 1 image 1 video
   {"CUK-1wjqqsP", 2669222102324390671},
   // 2 image 0 video
   {"CXzGW6RPNmy", 2734557361417804210},
}

func TestMedia(t *testing.T) {
   for _, test := range tests {
      med, err := newMedia(test.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", med)
      time.Sleep(time.Second)
   }
}

func TestID(t *testing.T) {
   for _, test := range tests {
      id, err := getID(test.shortcode)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}

