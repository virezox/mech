package tiktok

import (
   "fmt"
   "testing"
)

// tiktok.com/@kaimanoff/video/6896523341402737921
const id = 6896523341402737921

func TestDetail(t *testing.T) {
   det, err := NewDetail(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(det)
}

var addrs = []string{
   "https://vm.tiktok.com/ZMLesneqK",
   "https://www.tiktok.com/@eddysayi/video/7054218882072055046?_d=secCgwIARCbD",
}

func TestID(t *testing.T) {
   for _, addr := range addrs {
      id, err := AwemeID(addr)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}

