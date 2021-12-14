package tiktok

import (
   "fmt"
   "testing"
)

const addr = "https://www.tiktok.com/@aamora_3mk/video/7028702876205632773"

func TestData(t *testing.T) {
   item, err := NewItemStruct(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
