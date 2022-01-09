package pandora

import (
   "bytes"
   "encoding/hex"
   "testing"
)

const helloEnc = "7be654d97cc31582815d713a9d0c64ab"

var helloDec = []byte("hello world")

func TestDecrypt(t *testing.T) {
   enc, err := hex.DecodeString(helloEnc)
   if err != nil {
      t.Fatal(err)
   }
   dec := Cipher{enc}.Decrypt().Unpad()
   if !bytes.Equal(dec.Bytes, helloDec) {
      t.Fatal(dec)
   }
}
