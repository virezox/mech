package pandora

import (
   "bytes"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   LogLevel = 1
   user, err := part.userLogin("srpen6@gmail.com", password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", user)
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      t.Fatal("userAuthToken", tLen)
   }
}
