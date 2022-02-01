package instagram

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var shortcodes = []string{
   "CU9ett-rP7I",
   "CUDJ4YhpF0Z",
   "CUK-1wjqqsP",
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

func TestMedia(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := OpenLogin(cache + "/mech/instagram.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, shortcode := range shortcodes {
      med, err := login.Media(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(shortcode)
      for _, addr := range med.URLs() {
         fmt.Println("-", addr)
      }
      time.Sleep(time.Second)
   }
}
