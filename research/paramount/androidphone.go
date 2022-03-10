package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
)

func newToken() (string, error) {
   src := []byte{'|'}
   src = append(src, tv_secret...)
   key, err := hex.DecodeString(aes_key)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   iv := []byte("0123456789ABCDEF")
   src = pad(src)
   cipher.NewCBCEncrypter(block, iv).CryptBlocks(src, src)
   dst := []byte{0, 16}
   dst = append(dst, iv...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

const (
   aes_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"
   tv_secret = "6c70b33080758409"
)

func pad(c []byte) []byte {
   cLen := aes.BlockSize - len(c) % aes.BlockSize
   for high := byte(cLen); cLen >= 1; cLen-- {
      c = append(c, high)
   }
   return c
}
