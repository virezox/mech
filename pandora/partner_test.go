package pandora

import (
   "encoding/hex"
   "fmt"
   "testing"
)

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString("7be654d97cc3158278230e8df327865b")
   if err != nil {
      t.Fatal(err)
   }
   dec, err := Decrypt(enc)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", dec)
}
