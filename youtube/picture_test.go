package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

const id = "UpNXI3_ctAc"

func TestPicture(t *testing.T) {
   for _, p := range Pictures {
      addr := p.Address(id)
      fmt.Println("Head", addr)
      res, err := http.Head(addr)
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(100 * time.Millisecond)
   }
}
