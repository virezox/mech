package mech_test

import (
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "testing"
)

func TestEncoding(t *testing.T) {
   r, err := http.NewRequest("HEAD", "https://github.com/manifest.json", nil)
   if err != nil {
      t.Fatal(err)
   }
   c, err := mech.NewContent(r)
   if err != nil {
      t.Fatal(err)
   }
   for key, val := range c {
      fmt.Printf("%v %+v\n", key, val)
   }
}
