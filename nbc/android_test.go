package nbc

import (
   "fmt"
   "testing"
)

func TestAndroid(t *testing.T) {
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
