package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "fmt"
   "os"
   "encoding/json"
   "encoding/base64"
   "github.com/dchest/cmac"
   "github.com/89z/format/protobuf"
   "google.golang.org/protobuf/proto"
   wv "github.com/89z/mech/research/widevine"
)

var (
   _ = base64.StdEncoding
   _ = fmt.Print
   _ = json.NewEncoder
   _ = os.Stdout
   _ = protobuf.Unmarshal
)

func (c *decryptionModule) getLicenseKeys(licenseRequest, licenseResponse []byte) ([]licenseKey, error) {
   var license wv.SignedLicense
   err := proto.Unmarshal(licenseResponse, &license)
   if err != nil {
      return nil, err
   }
   var licenseRequestParsed wv.SignedLicenseRequest
   if err := proto.Unmarshal(licenseRequest, &licenseRequestParsed); err != nil {
      return nil, err
   }
   licenseRequestMsg, err := proto.Marshal(licenseRequestParsed.Msg)
   if err != nil {
      return nil, err
   }
   /////////////////////////////////////////////////////////////////////////////
   license2, err := protobuf.Unmarshal(licenseResponse)
   if err != nil {
      return nil, err
   }
   licenseRequestParsed2, err := protobuf.Unmarshal(licenseRequest)
   if err != nil {
      return nil, err
   }
   licenseRequestMsg2 := licenseRequestParsed2.Get(2).Marshal()
   /////////////////////////////////////////////////////////////////////////////
   fmt.Println(len(licenseRequestMsg), len(licenseRequestMsg2))
   /////////////////////////////////////////////////////////////////////////////
   sessionKey, err := rsa.DecryptOAEP(
      sha1.New(), nil, c.privateKey, license.SessionKey, nil,
   )
   if err != nil {
      return nil, err
   }
   /////////////////////////////////////////////////////////////////////////////
   key, err := license2.GetBytes(4)
   if err != nil {
      return nil, err
   }
   sessionKey2, err := rsa.DecryptOAEP(sha1.New(), nil, c.privateKey, key, nil)
   if err != nil {
      return nil, err
   }
   /////////////////////////////////////////////////////////////////////////////
   fmt.Println(len(license.SessionKey), len(key))
   /////////////////////////////////////////////////////////////////////////////
   sessionKeyBlock, err := aes.NewCipher(sessionKey)
   if err != nil {
      return nil, err
   }
   encryptionKey := []byte{
      1, 'E', 'N', 'C', 'R', 'Y', 'P', 'T', 'I', 'O', 'N', 0,
   }
   encryptionKey = append(encryptionKey, licenseRequestMsg...)
   encryptionKey = append(encryptionKey, []byte{0, 0, 0, 0x80}...)
   mac, err := cmac.New(sessionKeyBlock)
   if err != nil {
      return nil, err
   }
   mac.Write(encryptionKey)
   encryptionKeyCipher, err := aes.NewCipher(mac.Sum(nil))
   if err != nil {
      return nil, err
   }
   /////////////////////////////////////////////////////////////////////////////
   keyBlock2, err := aes.NewCipher(sessionKey2)
   if err != nil {
      return nil, err
   }
   encryptionKey2 := []byte{
      1, 'E', 'N', 'C', 'R', 'Y', 'P', 'T', 'I', 'O', 'N', 0,
   }
   encryptionKey2 = append(encryptionKey2, licenseRequestMsg2...)
   encryptionKey2 = append(encryptionKey2, []byte{0, 0, 0, 0x80}...)
   mac2, err := cmac.New(keyBlock2)
   if err != nil {
      return nil, err
   }
   mac2.Write(encryptionKey2)
   keyCipher2, err := aes.NewCipher(mac2.Sum(nil))
   if err != nil {
      return nil, err
   }
   /////////////////////////////////////////////////////////////////////////////
   fmt.Println(len(mac.Sum(nil)), len(mac2.Sum(nil)))
   /////////////////////////////////////////////////////////////////////////////
   var keys []licenseKey
   for _, key := range license.Msg.Key {
      decrypter := cipher.NewCBCDecrypter(encryptionKeyCipher, key.Iv)
      decryptedKey := make([]byte, len(key.Key))
      decrypter.CryptBlocks(decryptedKey, key.Key)
      keys = append(keys, licenseKey{
         Type:  *key.Type,
         Value: unpad(decryptedKey),
      })
   }
   var keys2 []licenseKey
   for _, con := range license2.Get(2).GetMessages(3) {
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
      decrypter := cipher.NewCBCDecrypter(keyCipher2, iv)
      decryptedKey := make([]byte, len(key))
      decrypter.CryptBlocks(decryptedKey, key)
      keys2 = append(keys2, licenseKey{
         Type2: uint64(typ),
         //Value: unpad(decryptedKey),
         Value: decryptedKey,
      })
   }
   fmt.Printf("%+v\n", keys2)
   return keys, nil
}

type licenseKey struct {
   Type  wv.License_KeyContainer_KeyType
   Type2  uint64
   Value []byte
}

type decryptionModule struct {
   clientID   []byte
   privateKey *rsa.PrivateKey
   //signedDeviceCertificate wv.SignedDeviceCertificate
   cencHeader struct {
      KeyId []byte "2"
   }
}

