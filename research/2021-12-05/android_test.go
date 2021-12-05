package nbc

import (
   "fmt"
   "testing"
)

func TestAndroid(t *testing.T) {
   vod, err := newAccessVOD(9000194212)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", vod)
}
