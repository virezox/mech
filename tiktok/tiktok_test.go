package tiktok

import (
   "fmt"
   "testing"
   "time"
)

var ids = []uint64{
   // tiktok.com/@aamora_3mk/video/7028702876205632773
   7028702876205632773,
   // tiktok.com/@elpanaarabe/video/7038818332270808325
   7038818332270808325,
}

func TestDetail(t *testing.T) {
   for _, id := range ids {
      det, err := NewAwemeDetail(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", det)
      time.Sleep(time.Second)
   }
}
