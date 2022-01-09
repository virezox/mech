package pandora

import (
   "encoding/hex"
   "golang.org/x/crypto/blowfish" //lint:ignore SA1019 reason
)

var blowfishKey = []byte("6#26FRL$ZWD")

type Cipher struct {
   Bytes []byte
}

func Decode(s string) *Cipher {
   buf, err := hex.DecodeString(s)
   if err != nil {
      return nil
   }
   return &Cipher{buf}
}

func (c Cipher) Decrypt() *Cipher {
   sLen := len(c.Bytes)
   if sLen < blowfish.BlockSize {
      return nil
   }
   block, err := blowfish.NewCipher(blowfishKey)
   if err != nil {
      return nil
   }
   for low := 0; low < sLen; low += blowfish.BlockSize {
      block.Decrypt(c.Bytes[low:], c.Bytes[low:])
   }
   return &c
}

func (c Cipher) Encode() string {
   return hex.EncodeToString(c.Bytes)
}

func (c Cipher) Encrypt() *Cipher {
   block, err := blowfish.NewCipher(blowfishKey)
   if err != nil {
      return nil
   }
   for low := 0; low < len(c.Bytes); low += blowfish.BlockSize {
      block.Encrypt(c.Bytes[low:], c.Bytes[low:])
   }
   return &c
}

func (c Cipher) Pad() Cipher {
   bLen := blowfish.BlockSize - len(c.Bytes) % blowfish.BlockSize
   for high := byte(bLen); bLen >= 1; bLen-- {
      c.Bytes = append(c.Bytes, high)
   }
   return c
}

func (c Cipher) Unpad() *Cipher {
   bLen := len(c.Bytes)
   if bLen == 0 {
      return nil
   }
   high := bLen - int(c.Bytes[bLen-1])
   if high <= -1 {
      return nil
   }
   c.Bytes = c.Bytes[:high]
   return &c
}
