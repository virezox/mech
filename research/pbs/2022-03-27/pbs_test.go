package pbs

import (
   "fmt"
   "testing"
)

func TestJSON(t *testing.T) {
   video, err := NewVideoBridge(3016754074)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
