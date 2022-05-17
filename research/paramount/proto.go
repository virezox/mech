package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "github.com/89z/format/protobuf"
   "github.com/chmike/cmac-go"
)

func (c *decryptionModule) getLicenseKeys(licenseRequest, licenseResponse []byte) ([]licenseKey, error) {
   license, err := protobuf.Unmarshal(licenseResponse)
   if err != nil {
      return nil, err
   }
   requestParsed, err := protobuf.Unmarshal(licenseRequest)
   if err != nil {
      return nil, err
   }
   requestMsg := requestParsed.Get(2).Marshal()
   cipherText, err := license.GetBytes(4)
   if err != nil {
      return nil, err
   }
   sessionKey, err := rsa.DecryptOAEP(
      sha1.New(), nil, c.privateKey, cipherText, nil,
   )
   if err != nil {
      return nil, err
   }
   hash, err := cmac.New(aes.NewCipher, sessionKey)
   if err != nil {
      return nil, err
   }
   var key []byte
   key = append(key, 1)
   key = append(key, "ENCRYPTION"...)
   key = append(key, 0)
   key = append(key, requestMsg...)
   key = append(key, 0, 0, 0, 0x80)
   hash.Write(key)
   block, err := aes.NewCipher(hash.Sum(nil))
   if err != nil {
      return nil, err
   }
   var keys []licenseKey
   for _, con := range license.Get(2).GetMessages(3) {
      iv, err := con.GetBytes(2)
      if err != nil {
         return nil, err
      }
      key, err := con.GetBytes(3)
      if err != nil {
         return nil, err
      }
      typ, err := con.GetVarint(4)
      if err != nil {
         return nil, err
      }
      decrypter := cipher.NewCBCDecrypter(block, iv)
      decryptedKey := make([]byte, len(key))
      decrypter.CryptBlocks(decryptedKey, key)
      keys = append(keys, licenseKey{
         Type2: uint64(typ),
         Value: unpad(decryptedKey),
      })
   }
   return keys, nil
}

type licenseKey struct {
   Type2  uint64
   Value []byte
}

type decryptionModule struct {
   clientID   []byte
   privateKey *rsa.PrivateKey
   cencHeader struct {
      KeyId []byte "2"
   }
}
