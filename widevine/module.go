package widevine

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

func (c Client) Module() (*Module, error) {
   key_id, err := c.KeyId()
   if err != nil {
      return nil, err
   }
   block, _ := pem.Decode(c.PrivateKey)
   var mod Module
   mod.PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   mod.licenseRequest, err = protobuf.Message{
      1: protobuf.Bytes{Raw: c.Id},
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes{Raw: key_id},
            },
         },
      },
   }.MarshalBinary()
   if err != nil {
      return nil, err
   }
   return &mod, nil
}

type Module struct {
   *rsa.PrivateKey
   licenseRequest []byte
}

func (m Module) Marshal() ([]byte, error) {
   digest := sha1.Sum(m.licenseRequest)
   signature, err := rsa.SignPSS(
      nopSource{},
      m.PrivateKey,
      crypto.SHA1,
      digest[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   signedRequest := protobuf.Message{
      2: protobuf.Bytes{Raw: m.licenseRequest},
      3: protobuf.Bytes{Raw: signature},
   }
   return signedRequest.MarshalBinary()
}

func (m Module) Unmarshal(response []byte) (Contents, error) {
   // key
   signedResponse := make(protobuf.Message)
   err := signedResponse.UnmarshalBinary(response)
   if err != nil {
      return nil, err
   }
   sessionKey, err := signedResponse.GetBytes(4)
   if err != nil {
      return nil, err
   }
   key, err := rsa.DecryptOAEP(sha1.New(), nil, m.PrivateKey, sessionKey, nil)
   if err != nil {
      return nil, err
   }
   // message
   var message []byte
   message = append(message, 1)
   message = append(message, "ENCRYPTION"...)
   message = append(message, 0)
   message = append(message, m.licenseRequest...)
   message = append(message, 0, 0, 0, 0x80)
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
   var cons Contents
   // .Msg.Key
   for _, message := range signedResponse.Get(2).GetMessages(3) {
      var con Content
      iv, err := message.GetBytes(2)
      if err != nil {
         return nil, err
      }
      con.Key, err = message.GetBytes(3)
      if err != nil {
         return nil, err
      }
      con.Type, err = message.GetVarint(4)
      if err != nil {
         return nil, err
      }
      cipher.NewCBCDecrypter(block, iv).CryptBlocks(con.Key, con.Key)
      con.Key = unpad(con.Key)
      cons = append(cons, con)
   }
   return cons, nil
}
