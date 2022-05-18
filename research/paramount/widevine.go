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
   "io"
)

func unpad(buf []byte) []byte {
   if len(buf) >= 1 {
      pad := buf[len(buf)-1]
      if len(buf) >= int(pad) {
         buf = buf[:len(buf)-int(pad)]
      }
   }
   return buf
}

type KeyContainer struct {
   Key []byte
   Type uint64
}

type Module struct {
   *rsa.PrivateKey
   licenseRequest []byte
}

func NewModule(privateKey, clientID, kID []byte) (*Module, error) {
   var (
      err error
      mod Module
   )
   // PrivateKey
   block, _ := pem.Decode(privateKey)
   mod.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   // licenseRequest
   mod.licenseRequest = protobuf.Message{
      1: protobuf.Bytes(clientID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(kID),
            },
         },
      },
   }.Marshal()
   return &mod, nil
}

func (m *Module) Keys(licenseResponse io.Reader) ([]KeyContainer, error) {
   // message
   var message []byte
   message = append(message, 1)
   message = append(message, "ENCRYPTION"...)
   message = append(message, 0)
   message = append(message, m.licenseRequest...)
   message = append(message, 0, 0, 0, 0x80)
   // key
   signedLicense, err := protobuf.Decode(licenseResponse)
   if err != nil {
      return nil, err
   }
   sessionKey, err := signedLicense.GetBytes(4)
   if err != nil {
      return nil, err
   }
   key, err := rsa.DecryptOAEP(sha1.New(), nil, m.PrivateKey, sessionKey, nil)
   if err != nil {
      return nil, err
   }
   // CMAC
   mac, err := cmac.New(aes.NewCipher, key)
   if err != nil {
      return nil, err
   }
   mac.Write(message)
   block, err := aes.NewCipher(mac.Sum(nil))
   if err != nil {
      return nil, err
   }
   var containers []KeyContainer
   // .Msg.Key
   for _, message := range signedLicense.Get(2).GetMessages(3) {
      iv, err := message.GetBytes(2)
      if err != nil {
         return nil, err
      }
      key, err := message.GetBytes(3)
      if err != nil {
         return nil, err
      }
      typ, err := message.GetVarint(4)
      if err != nil {
         return nil, err
      }
      cipher.NewCBCDecrypter(block, iv).CryptBlocks(key, key)
      var container KeyContainer
      container.Key = unpad(key)
      container.Type = uint64(typ)
      containers = append(containers, container)
   }
   return containers, nil
}

func (c *Module) SignedLicenseRequest() ([]byte, error) {
   digest := sha1.Sum(c.licenseRequest)
   signature, err := rsa.SignPSS(
      nopSource{},
      c.PrivateKey,
      crypto.SHA1,
      digest[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   signedLicenseRequest := protobuf.Message{
      2: protobuf.Bytes(c.licenseRequest),
      3: protobuf.Bytes(signature),
   }
   return signedLicenseRequest.Marshal(), nil
}

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
}
