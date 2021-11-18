package bleep

import (
   "fmt"
   "testing"
)

func TestBleep(t *testing.T) {
   id, err := ReleaseID("https://bleep.com/release/8728-four-tet-pause")
   if err != nil {
      t.Fatal(err)
   }
   rel, err := Release(id)
   if err != nil {
      t.Fatal(err)
   }
   for _, track := range rel {
      addr, err := track.Resolve()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(addr)
      break
   }
}
