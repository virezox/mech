package tiktok

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

const video = "https://www.tiktok.com/@aamora_3mk/video/7028702876205632773"

func TestData(t *testing.T) {
   mech.Verbose = true
   d, err := NewData(video)
   if err != nil {
      t.Fatal(err)
   }
   r, err := GetVideo(d.PlayAddr())
   if err != nil {
      t.Fatal(err)
   }
   defer r.Body.Close()
   fmt.Printf("%+v\n", r)
}
