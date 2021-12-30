package pandora

import (
   "bytes"
   "encoding/hex"
   "fmt"
   "net/url"
   "testing"
)

const helloEnc = "7be654d97cc31582815d713a9d0c64ab"

var helloDec = []byte("hello world")

func TestValues(t *testing.T) {
   val := url.Values{
      "one": {},
      "two": {"three"},
   }.Encode()
   fmt.Println(val)
}

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

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(helloEnc)
   if err != nil {
      t.Fatal(err)
   }
   dec, err := Decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(dec, helloDec) {
      t.Fatal(dec)
   }
}
