package instagram

import (
   "fmt"
   "testing"
   "time"
)

const shortcode = "CT-cnxGhvvO"

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
