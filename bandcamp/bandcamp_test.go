package bandcamp

import (
   "fmt"
   "testing"
)

const art_ID = 3809045440

func Test_Image(t *testing.T) {
   for _, img := range Images {
      addr := img.URL(art_ID)
      fmt.Println(addr)
   }
}
