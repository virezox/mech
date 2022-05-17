package paramount

import (
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/pem"
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

func (c *decryptionModule) getLicenseRequest() ([]byte, error) {
   msg := protobuf.Message{
      1: protobuf.Bytes(c.clientID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(c.cencHeader.KeyId),
            },
         },
      },
   }
   hash := sha1.Sum(msg.Marshal())
   signature, err := rsa.SignPSS(
      nopSource{},
      c.privateKey,
      crypto.SHA1,
      hash[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   licenseRequest := protobuf.Message{
      2: msg,
      3: protobuf.Bytes(signature),
   }
   return licenseRequest.Marshal(), nil
}

func newCDM(privateKey, clientID, initData []byte) (*decryptionModule, error) {
   block, _ := pem.Decode(privateKey)
   keyParsed, err := x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      // if PCKS1 doesn't work, try PCKS8
      pcks8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
      if err != nil {
         return nil, err
      }
      keyParsed = pcks8Key.(*rsa.PrivateKey)
   }
   var dec decryptionModule
   dec.clientID = clientID
   dec.privateKey = keyParsed
   mes, err := protobuf.Unmarshal(initData[32:])
   if err != nil {
      return nil, err
   }
   dec.cencHeader.KeyId, err = mes.GetBytes(2)
   if err != nil {
      return nil, err
   }
   return &dec, nil
}

