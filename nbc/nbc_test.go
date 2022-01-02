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
   forms, err := vod.Manifest()
   if err != nil {
      t.Fatal(err)
   }
   for _, form := range forms {
      fmt.Println(form)
   }
}

func TestAndroidVideo(t *testing.T) {
   LogLevel = 1
   vid, err := NewVideo(guid)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}
