package paramount

import (
   "crypto"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/pem"
   "github.com/89z/format/protobuf"
)

func newModule(privateKey, clientID, initData []byte) (*module, error) {
   var dec module
   dec.clientID = clientID
   mes, err := protobuf.Unmarshal(initData[32:])
   if err != nil {
      return nil, err
   }
   block, _ := pem.Decode(privateKey)
   dec.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   dec.KeyId, err = mes.GetBytes(2)
   if err != nil {
      return nil, err
   }
   return &dec, nil
}

func (c *module) getLicenseRequest() ([]byte, error) {
   msg := protobuf.Message{
      1: protobuf.Bytes(c.clientID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(c.KeyId),
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
