package youtube

import (
   "fmt"
   "testing"
   "time"
)

func Test_OAuth(t *testing.T) {
   auth, err := New_OAuth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v
`, auth.Verification_URL, auth.User_Code)
   for range [9]bool{} {
      time.Sleep(9 * time.Second)
      change, err := auth.Exchange()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", change)
      if change.Access_Token != "" {
         break
      }
   }
}
