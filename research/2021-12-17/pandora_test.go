package pandora

import (
   "fmt"
   "os"
   "testing"
)

const partnerAuthToken = "VAr6MH7yWoDPh0qkHj682aHQ=="

func TestLogin(t *testing.T) {
   buf, err := os.ReadFile("fail.txt")
   if err != nil {
      t.Fatal(err)
   }
   LogLevel = 1
   user, err := newUserLogin(partnerAuthToken, buf)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", user)
   tLen := len(user.Result.UserAuthToken)
   if tLen != 58 {
      t.Fatal("userAuthToken", tLen)
   }
}
