package pandora

import (
   "fmt"
   "testing"
)

func TestPartner(t *testing.T) {
   LogLevel = 1
   part, err := newPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   tLen := len(part.Result.PartnerAuthToken)
   if tLen != 34 {
      t.Fatal("partnerAuthToken", tLen)
   }
}

func TestLogin(t *testing.T) {
   LogLevel = 1
   user, err := newUserLogin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", user)
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      t.Fatal("userAuthToken", tLen)
   }
}
