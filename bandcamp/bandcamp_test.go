package bandcamp

import (
   "fmt"
   "testing"
)

const artID = 3809045440

func TestImage(t *testing.T) {
   for _, img := range Images {
      addr := img.Format(artID)
      fmt.Println(addr)
   }
}
