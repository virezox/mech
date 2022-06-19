package widevine

import (
   "bytes"
   "encoding/base64"
   "testing"
)

const pssh = "AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAAAAAWgwC7YzAgICAgICBI88aJmwY="

func TestKey(t *testing.T) {
   key_a, err := base64.StdEncoding.DecodeString("AAAAABaDALtjMCAgICAgIA==")
   if err != nil {
      t.Fatal(err)
   }
   key_b, err := Client{Raw_PSSH: pssh}.Key_ID()
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(key_a, key_b) {
      t.Fatal(key_b)
   }
}
