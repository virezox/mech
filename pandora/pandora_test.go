package pandora

import (
   "testing"
)

func TestLogin(t *testing.T) {
   part, err := NewPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   tLen := len(part.Result.PartnerAuthToken)
   if tLen != 34 {
      t.Fatal(tLen)
   }
   user, err := part.UserLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   if tLen := len(user.Result.UserAuthToken); tLen != 58 {
      t.Fatal(tLen)
   }
}
