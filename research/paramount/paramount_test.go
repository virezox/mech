package paramount

import (
   "fmt"
   "testing"
)

const addr = "paramountplus.com/shows/star-trek-prodigy/video/3htV4fvVt4Z8gDZHqlzPOGLSMgcGc_vy/star-trek-prodigy-dreamcatcher"

func TestParamount(t *testing.T) {
   id := VideoID(addr)
   addr, err := Media(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(addr)
}
