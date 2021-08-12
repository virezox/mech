package soundcloud

import (
   "fmt"
   "testing"
)

func TestSoundCloud(t *testing.T) {
   var opt APIOptions
   api, err := New(opt)
   if err != nil {
      t.Fatal(err)
   }
   addr, err := api.GetDownloadURL(
      "https://soundcloud.com/bluewednesday/murmuration-feat-shopan",
      "progressive",
   )
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(addr)
}
