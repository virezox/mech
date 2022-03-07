package instagram

import (
   "fmt"
   "net/url"
   "os"
   "testing"
   "time"
)

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
      items, err := login.Items(test.shortcode)
      if err != nil {
         t.Fatal(err)
      }
      for _, item := range items {
         var index int
         for _, med := range item.Medias() {
            addrs, err := med.URLs()
            if err != nil {
               t.Fatal(err)
            }
            for _, address := range addrs {
               addr, err := url.Parse(address)
               if err != nil {
                  t.Fatal(err)
               }
               if addr.Path != test.paths[index] {
                  t.Fatalf("%q\n", addr.Path)
               }
               index++
            }
         }
      }
      time.Sleep(time.Second)
   }
}
