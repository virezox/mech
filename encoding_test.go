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
   for _, enc := range mech.AcceptEncoding {
      c, err := mech.NewContent(r, enc)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%v %+v\n", enc, c)
   }
}
