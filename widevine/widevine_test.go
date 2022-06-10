package widevine

import (
   "bytes"
   "encoding/base64"
   "testing"
)

const pssh = "AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="

func TestKey(t *testing.T) {
   keyA, err := base64.StdEncoding.DecodeString("AAAAABaDALtjMCAgICAgIA==")
   if err != nil {
      t.Fatal(err)
   }
   keyB, err := KeyID(pssh)
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(keyA, keyB) {
      t.Fatal(keyB)
   }
}
