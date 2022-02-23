package soundcloud

import (
   "fmt"
   "testing"
)

func TestUser(t *testing.T) {
   user, err := NewUserStream(692707328)
   if err != nil {
      t.Fatal(err)
   }
   for _, item := range user.Collection {
      fmt.Printf("%+v\n", item.Track)
   }
}
