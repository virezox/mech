package nbc

import (
   "fmt"
   "github.com/89z/mech"
   "testing"
)

const (
   // nbc.com/la-brea/video/pilot/9000194212
   res540 = 9000194212
   // nbc.com/the-blacklist/video/the-skinner/9000210182
   res1080 = 9000210182
)

func TestWeb(t *testing.T) {
   mech.Verbose = true
   forms, err := Media(res540)
   if err != nil {
      t.Fatal(err)
   }
   for _, form := range forms {
      fmt.Printf("%+v\n", form)
   }
}
