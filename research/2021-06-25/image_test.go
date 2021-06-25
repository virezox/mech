package youtube_test

import (
   "github.com/89z/mech/youtube"
   "net/http"
   "testing"
   "time"
)

const id = "UpNXI3_ctAc"

func TestImage(t *testing.T) {
   for _, img := range youtube.Images {
      addr := img.Address(id)
      println("Head", addr)
      r, err := http.Head(addr)
      if err != nil {
         t.Fatal(err)
      }
      if r.StatusCode != http.StatusOK {
         t.Fatal(r.Status)
      }
      time.Sleep(100 * time.Millisecond)
   }
}
