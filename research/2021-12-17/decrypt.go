package main

import (
   "bytes"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "golang.org/x/crypto/blowfish"
   "os"
)

var key = []byte("6#26FRL$ZWD")

func decrypt(src []byte) ([]byte, error) {
   dst := make([]byte, len(src))
   blow, err := blowfish.NewCipher(key)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(src); low += blowfish.BlockSize {
      blow.Decrypt(dst[low:], src[low:])
   }
   pad := dst[len(dst)-1]
   return dst[:len(dst)-int(pad)], nil
}

func main() {
   if len(os.Args) != 2 {
      fmt.Println("decrypt [string]")
      return
   }
   src := os.Args[1]
   enc, err := hex.DecodeString(src)
   if err != nil {
      panic(err)
   }
   dec, err := decrypt(enc)
   if err != nil {
      panic(err)
   }
   buf := new(bytes.Buffer)
   json.Indent(buf, dec, "", " ")
   os.Stdout.ReadFrom(buf)
}
