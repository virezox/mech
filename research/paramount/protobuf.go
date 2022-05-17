package paramount

import (
   "crypto"
   "crypto/rsa"
   "crypto/sha1"
   "github.com/89z/format/protobuf"
)

func (c *module) getLicenseRequest() ([]byte, error) {
   msg := protobuf.Message{
      1: protobuf.Bytes(c.clientID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(c.keyID),
            },
         },
      },
   }
   hash := sha1.Sum(msg.Marshal())
   signature, err := rsa.SignPSS(
      nopSource{},
      c.PrivateKey,
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
