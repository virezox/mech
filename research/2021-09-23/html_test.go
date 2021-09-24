package instagram

import (
   "fmt"
   "testing"
   "time"
)

func TestHtmlEmbed(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := htmlEmbed(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestHtmlP(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := htmlP(shortcode)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
