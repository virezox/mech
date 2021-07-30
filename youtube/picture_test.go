package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "net/http"
   "testing"
   "time"
)

const id = "UpNXI3_ctAc"

func TestImage(t *testing.T) {
   for _, p := range youtube.Pictures {
      addr := p.Address(id)
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

func TestFilter(t *testing.T) {
   pix := youtube.Pictures.Filter(func(i youtube.Picture) bool {
      return i.Height < 720
   })
   pix.Sort()
   for _, p := range pix {
      fmt.Println(p)
   }
}
