package nbc

import (
   "fmt"
   "testing"
)

// nbc.com/saturday-night-live/video/january-22-will-forte/9000199367
const guid = 9000199367

func TestAndroidVOD(t *testing.T) {
   vod, err := NewAccessVOD(guid)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vod)
}
