package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "fmt"
   "github.com/dchest/cmac"
   "google.golang.org/protobuf/proto"
   wv "github.com/89z/mech/research/widevine"
   "github.com/89z/format/protobuf"
)

var (
   _ = fmt.Print
   _ = protobuf.Unmarshal
)

// Retrieves the keys from the license response data.  These keys can be
// used to decrypt the DASH-MP4.
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
   sessionKey, err := rsa.DecryptOAEP(
      sha1.New(), nil, c.privateKey, license.SessionKey, nil,
   )
   if err != nil {
      return nil, err
   }
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
   var keys []licenseKey
   for _, key := range license.Msg.Key {
      decrypter := cipher.NewCBCDecrypter(encryptionKeyCipher, key.Iv)
      decryptedKey := make([]byte, len(key.Key))
      decrypter.CryptBlocks(decryptedKey, key.Key)
      keys = append(keys, licenseKey{
         ID:    key.Id,
         Type:  *key.Type,
         Value: unpad(decryptedKey),
      })
   }
   return keys, nil
}

type decryptionModule struct {
   clientID   []byte
   privateKey *rsa.PrivateKey
   signedDeviceCertificate wv.SignedDeviceCertificate
   cencHeader struct {
      KeyId []byte "2"
   }
}

type licenseKey struct {
   ID    []byte
   Type  wv.License_KeyContainer_KeyType
   Value []byte
}
