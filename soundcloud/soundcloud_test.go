package soundcloud

import (
   "fmt"
   "testing"
)

func TestSoundCloud(t *testing.T) {
   api, err := NewClient()
   if err != nil {
      t.Fatal(err)
   }
   addr, err := api.GetDownloadURL(
      "https://soundcloud.com/bluewednesday/murmuration-feat-shopan",
   )
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(addr)
}
