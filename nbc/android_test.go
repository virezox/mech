package nbc

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

func TestAndroid(t *testing.T) {
   mech.LogLevel = 2
   vod, err := NewAccessVOD(9000194212)
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
