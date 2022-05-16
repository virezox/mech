package paramount

import (
   "bytes"
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "fmt"
   "github.com/dchest/cmac"
   "google.golang.org/protobuf/proto"
   "time"
   wv "github.com/89z/mech/research/widevine"
)

var _ = fmt.Print

// Generates the license request data.  This is sent to the license server via
// HTTP POST and the server in turn returns the license response.
func (c *decryptionModule) getLicenseRequest() ([]byte, error) {
   var licenseRequest wv.SignedLicenseRequest
   licenseRequest.Msg = new(wv.LicenseRequest)
   licenseRequest.Msg.ContentId = new(wv.LicenseRequest_ContentIdentification)
   licenseRequest.Msg.ContentId.CencId = new(wv.LicenseRequest_ContentIdentification_CENC)
   licenseRequest.Msg.ContentId.CencId.Pssh = &wv.WidevineCencHeader{
      KeyId: [][]byte{c.cencHeader.KeyId},
   }
   // this is probably really bad for the GC but protobuf uses pointers for
   // optional fields so it is necessary and this is not a long running program
   licenseRequest.Msg.ContentId.CencId.RequestId = c.sessionID[:]
   {
      v := uint32(time.Now().Unix())
      licenseRequest.Msg.RequestTime = &v
   }
   {
      v := wv.ProtocolVersion_CURRENT
      licenseRequest.Msg.ProtocolVersion = &v
   }
   if c.privacyMode {
      pad := func(data []byte, blockSize int) []byte {
         padlen := blockSize - (len(data) % blockSize)
         if padlen == 0 {
            padlen = blockSize
         }
         return append(data, bytes.Repeat([]byte{byte(padlen)}, padlen)...)
      }
      const blockSize = 16
      var (
         cidIV [blockSize]byte
         cidKey [blockSize]byte
      )
      block, err := aes.NewCipher(cidKey[:])
      if err != nil {
         return nil, err
      }
      paddedClientID := pad(c.clientID, blockSize)
      encryptedClientID := make([]byte, len(paddedClientID))
      cipher.NewCBCEncrypter(block, cidIV[:]).CryptBlocks(encryptedClientID, paddedClientID)
      servicePublicKey, err := x509.ParsePKCS1PublicKey(c.signedDeviceCertificate.XDeviceCertificate.PublicKey)
      if err != nil {
         return nil, err
      }
      encryptedCIDKey, err := rsa.EncryptOAEP(
         sha1.New(), nil, servicePublicKey, cidKey[:], nil,
      )
      if err != nil {
         return nil, err
      }
      licenseRequest.Msg.EncryptedClientId = new(wv.EncryptedClientIdentification)
      {
         v := string(c.signedDeviceCertificate.XDeviceCertificate.ServiceId)
         licenseRequest.Msg.EncryptedClientId.ServiceId = &v
      }
      licenseRequest.Msg.EncryptedClientId.ServiceCertificateSerialNumber = c.signedDeviceCertificate.XDeviceCertificate.SerialNumber
      licenseRequest.Msg.EncryptedClientId.EncryptedClientId = encryptedClientID
      licenseRequest.Msg.EncryptedClientId.EncryptedClientIdIv = cidIV[:]
      licenseRequest.Msg.EncryptedClientId.EncryptedPrivacyKey = encryptedCIDKey
   } else {
      licenseRequest.Msg.ClientId = new(wv.ClientIdentification)
      if err := proto.Unmarshal(c.clientID, licenseRequest.Msg.ClientId); err != nil {
         return nil, err
      }
   }
   {
      data, err := proto.Marshal(licenseRequest.Msg)
      if err != nil {
         return nil, err
      }
      hash := sha1.Sum(data)
      if licenseRequest.Signature, err = rsa.SignPSS(
         nopSource{}, c.privateKey, crypto.SHA1, hash[:],
         &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
      ); err != nil {
         return nil, err
      }
   }
   return proto.Marshal(&licenseRequest)
}

type decryptionModule struct {
   clientID   []byte
   privacyMode             bool
   privateKey *rsa.PrivateKey
   sessionID  [32]byte
   signedDeviceCertificate wv.SignedDeviceCertificate
   cencHeader struct {
      KeyId []byte "2"
   }
}

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

type licenseKey struct {
   ID    []byte
   Type  wv.License_KeyContainer_KeyType
   Value []byte
}
