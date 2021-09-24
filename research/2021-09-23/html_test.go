package instagram

import (
   "fmt"
   "testing"
   "time"
)

func TestHtmlEmbed(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := htmlEmbed(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestHtmlP(t *testing.T) {
   for i := range [16]struct{}{} {
      fmt.Println(i)
      err := htmlP(id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
