package nbc

import (
   "fmt"
   "testing"
)

func TestNBC(t *testing.T) {
   res, err := vod()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   fmt.Printf("%+v\n", res)
}
