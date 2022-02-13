package imdb

import (
   "fmt"
   "testing"
)

const rgconst = "rg2774637312"

func TestCred(t *testing.T) {
   cred, err := NewCredential()
   if err != nil {
      t.Fatal(err)
   }
   gal, err := cred.Gallery(rgconst)
   if err != nil {
      t.Fatal(err)
   }
   for _, img := range gal.Images {
      fmt.Printf("%+v\n", img)
   }
}
