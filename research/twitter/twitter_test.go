package twitter

import (
   "fmt"
   "testing"
)

func TestTwitter(t *testing.T) {
   guest, err := NewGuest()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := guest.xauth(identifier, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
