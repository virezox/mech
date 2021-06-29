package youtube_test

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "math"
   "net/http"
   "sort"
   "testing"
   "time"
)

const id = "UpNXI3_ctAc"

func TestImage(t *testing.T) {
   return
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

type iFunc func(a, b youtube.Image) bool

var iFuncs = []iFunc{
   func(a, b youtube.Image) bool {
      return math.Copysign(1, a.Height-720) < math.Copysign(1, b.Height-720)
   },
   func(a, b youtube.Image) bool {
      return math.Abs(a.Height-720) < math.Abs(b.Height-720)
   },
   func(a, b youtube.Image) bool {
      return a.Frame < b.Frame
   },
   func(a, b youtube.Image) bool {
      return a.Format < b.Format
   },
}

func TestSort(t *testing.T) {
   sort.SliceStable(youtube.Images, func(a, b int) bool {
      ha, hb := youtube.Images[a], youtube.Images[b]
      for _, fn := range iFuncs {
         if fn(ha, hb) {
            return true
         }
         if fn(hb, ha) {
            break
         }
      }
      return false
   })
   for _, img := range youtube.Images {
      fmt.Println(img)
   }
}
