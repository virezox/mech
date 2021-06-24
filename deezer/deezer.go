// Deezer
package deezer

import (
   "crypto/cipher"
   "crypto/md5"
   "encoding/hex"
   "golang.org/x/crypto/blowfish"
)

const GatewayWWW = "http://www.deezer.com/ajax/gw-light.php"

const (
   FLAC = '9'
   MP3_320 = '3'
)

var (
   iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
   keyAES = []byte("jo6aey6haid2Teih")
   keyBlowfish = []byte("g4el58wc0zvf9na1")
)

// Given SNG_ID and byte slice, decrypt byte slice in place.
func Decrypt(sngID string, data []byte) error {
   hash := md5Hash(sngID)
   for n := range keyBlowfish {
      keyBlowfish[n] ^= hash[n] ^ hash[n + len(keyBlowfish)]
   }
   block, err := blowfish.NewCipher(keyBlowfish)
   if err != nil {
      return err
   }
   size := 2048
   for pos := 0; len(data) - pos >= size; pos += size {
      if pos / size % 3 == 0 {
         text := data[pos : pos + size]
         cipher.NewCBCDecrypter(block, iv).CryptBlocks(text, text)
      }
   }
   return nil
}

func md5Hash(s string) string {
   sum := md5.Sum([]byte(s))
   return hex.EncodeToString(sum[:])
}
