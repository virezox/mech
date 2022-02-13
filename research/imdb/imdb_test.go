package imdb

import (
   "fmt"
   "testing"
)

func TestIMDb(t *testing.T) {
   cred, err := newCredentials()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", cred)
}
