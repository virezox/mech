package nbc

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

// nbc.com/la-brea/video/pilot/9000194212
const guid = 9000194212

func TestAndroidVideo(t *testing.T) {
   vid, err := NewVideo(guid)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vid)
}

func TestAndroidVOD(t *testing.T) {
   mech.LogLevel = 2
   vod, err := NewAccessVOD(guid)
   if err != nil {
      t.Fatal(err)
   }
   forms, err := vod.Manifest()
   if err != nil {
      t.Fatal(err)
   }
   for _, form := range forms {
      delete(form, "URI")
      fmt.Println(form)
   }
}
