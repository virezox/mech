package instagram

import (
   "fmt"
   "os"
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
