package pandora

import (
   "encoding/hex"
   "fmt"
   "testing"
)

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(loginEnc)
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
   enc, err := encrypt(loginDec)
   if err != nil {
      t.Fatal(err)
   }
   str := hex.EncodeToString(enc)
   fmt.Println(str)
}

func TestLogin(t *testing.T) {
   login, err := newPartnerLogin()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login)
}
