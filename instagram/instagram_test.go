package instagram

import (
   "fmt"
   "os"
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
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      med, err := login.Media(test.shortcode)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(test.shortcode)
      for _, addr := range med.URLs() {
         fmt.Println("-", addr)
      }
      time.Sleep(time.Second)
   }
}

func TestWrite(t *testing.T) {
   login, err := NewLogin("srpen6", password)
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Create("instagram.json"); err != nil {
      t.Fatal(err)
   }
}
