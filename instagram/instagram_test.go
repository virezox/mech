package instagram

import (
   "fmt"
   "testing"
   "time"
)

var shortcodes = []string{
   "CU9ett-rP7I", // "https://s" "" nil
   "CUDJ4YhpF0Z", // "https://s" "https://s" nil
   "CUK-1wjqqsP", // "https://s" "" "https://s" "https://s"
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
   for _, shortcode := range shortcodes {
      med, err := newMedia(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      short := med.Data.Shortcode_Media
      fmt.Printf("%v %.9q %.9q", shortcode, short.Display_URL, short.Video_URL)
      if short.Edge_Sidecar_To_Children != nil {
         node := short.Edge_Sidecar_To_Children.Edges[1].Node
         fmt.Printf(" %.9q %.9q\n", node.Display_URL, node.Video_URL)
      } else {
         fmt.Println(" nil")
      }
      time.Sleep(time.Second)
   }
}
