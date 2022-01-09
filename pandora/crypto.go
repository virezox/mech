package pandora

import (
   "encoding/hex"
   "encoding/json"
   "github.com/89z/format"
   "golang.org/x/crypto/blowfish" //lint:ignore SA1019 reason
)

var blowfishKey = []byte("6#26FRL$ZWD")

func Decrypt(src []byte) ([]byte, error) {
   sLen := len(src)
   if sLen < blowfish.BlockSize {
      return nil, format.InvalidSlice{blowfish.BlockSize-1, sLen}
   }
   dst := make([]byte, sLen)
   block, err := blowfish.NewCipher(blowfishKey)
   if err != nil {
      return nil, err
   }
   for low := 0; low < sLen; low += blowfish.BlockSize {
      block.Decrypt(dst[low:], src[low:])
   }
   return unpad(dst)
}

func Encrypt(src []byte) ([]byte, error) {
   src = pad(src)
   dst := make([]byte, len(src))
   block, err := blowfish.NewCipher(blowfishKey)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(src); low += blowfish.BlockSize {
      block.Encrypt(dst[low:], src[low:])
   }
   return dst, nil
}

func hexEncode(val interface{}) (string, error) {
   body, err := json.Marshal(val)
   if err != nil {
      return "", err
   }
   buf, err := Encrypt(body)
   if err != nil {
      return "", err
   }
   return hex.EncodeToString(buf), nil
}

func pad(src []byte) []byte {
   sLen := blowfish.BlockSize - len(src) % blowfish.BlockSize
   for high := byte(sLen); sLen >= 1; sLen-- {
      src = append(src, high)
   }
   return src
}

func unpad(src []byte) ([]byte, error) {
   sLen := len(src)
   if sLen == 0 {
      return nil, format.InvalidSlice{}
   }
   tLen := src[sLen-1]
   high := sLen - int(tLen)
   if high <= -1 {
      return nil, format.InvalidSlice{high, sLen}
   }
   return src[:high], nil
}
