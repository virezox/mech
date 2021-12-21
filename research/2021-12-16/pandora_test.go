package pandora

import (
   "encoding/hex"
   "fmt"
   "testing"
)

func TestLogin(t *testing.T) {
   part, err := newPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", part)
   user, err := part.userLogin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", user)
}

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(userLoginEnc)
   if err != nil {
      t.Fatal(err)
   }
   dec, err := decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", dec)
}

func TestEncrypt(t *testing.T) {
   enc, err := encrypt(userLoginDec)
   if err != nil {
      t.Fatal(err)
   }
   str := hex.EncodeToString(enc)
   fmt.Println(str)
}
