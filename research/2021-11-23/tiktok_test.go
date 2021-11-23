package tiktok

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
   "time"
)

const video = "https://www.tiktok.com/@aamora_3mk/video/7028702876205632773"

func TestData(t *testing.T) {
   mech.Verbose = true
   for range [9]struct{}{} {
      d, err := newData(video)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(d.playAddr())
      time.Sleep(time.Second)
   }
}
