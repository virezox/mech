package twitter

import (
   "fmt"
   "testing"
)

// twitter.com/PrimeTobirama/status/1577759019724333058
const id = 1577759019724333058

func Test_Status(t *testing.T) {
   Client.Log_Level = 2
   g, err := New_Guest()
   if err != nil {
      t.Fatal(err)
   }
   s, err := g.Status(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(s)
}
