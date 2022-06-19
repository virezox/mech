package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestOauth(t *testing.T) {
   auth, err := NewOauth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v
`, auth.Verification_URL, auth.User_Code)
   for range [9]struct{}{} {
      time.Sleep(9 * time.Second)
      exc, err := auth.Exchange()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", exc)
      if exc.Access_Token != "" {
         break
      }
   }
}
