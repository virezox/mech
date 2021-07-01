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
   for _, img := range youtube.AdaptiveImages {
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

func TestFilter(t *testing.T) {
   imgs := youtube.AdaptiveImages.Filter(func(i youtube.Image) bool {
      return i.Height < 720
   })
   imgs.Sort()
   for _, img := range imgs {
      fmt.Println(img)
   }
}
