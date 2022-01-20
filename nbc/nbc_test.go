package nbc

import (
   "fmt"
   "testing"
)

// nbc.com/la-brea/video/pilot/9000194212
const guid = 9000194212

func TestAndroidVOD(t *testing.T) {
   vod, err := NewAccessVOD(guid)
   if err != nil {
      t.Fatal(err)
   }
   streams, err := vod.Streams()
   if err != nil {
      t.Fatal(err)
   }
   for _, stream := range streams {
      fmt.Println(stream)
   }
}

func TestAndroidVideo(t *testing.T) {
   vid, err := NewVideo(guid)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
