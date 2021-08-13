package soundcloud

import (
   "fmt"
   "testing"
)

func TestSoundCloud(t *testing.T) {
   api, err := New()
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
