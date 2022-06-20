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
   key_id, err := c.Key_ID()
   if err != nil {
      return nil, err
   }
   block, _ := pem.Decode(c.Private_Key)
   var mod Module
   mod.private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   mod.licenseRequest = protobuf.Message{
      1: protobuf.Bytes{Raw: c.ID},
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes{Raw: key_id},
            },
         },
      },
   }.Marshal()
   return &mod, nil
}

type Module struct {
   licenseRequest []byte
   private_key *rsa.PrivateKey
}

func (m Module) Marshal() ([]byte, error) {
   digest := sha1.Sum(m.licenseRequest)
   signature, err := rsa.SignPSS(
      nopSource{},
      m.private_key,
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
   return signedRequest.Marshal(), nil
}

func (m Module) Unmarshal(response []byte) (Contents, error) {
   // key
   signed_response, err := protobuf.Unmarshal(response)
   if err != nil {
      return nil, err
   }
   sessionKey, err := signed_response.Get_Bytes(4)
   if err != nil {
      return nil, err
   }
   key, err := rsa.DecryptOAEP(sha1.New(), nil, m.private_key, sessionKey, nil)
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
   for _, message := range signed_response.Get(2).Get_Messages(3) {
      var con Content
      iv, err := message.Get_Bytes(2)
      if err != nil {
         return nil, err
      }
      con.Key, err = message.Get_Bytes(3)
      if err != nil {
         return nil, err
      }
      con.Type, err = message.Get_Varint(4)
      if err != nil {
         return nil, err
      }
      cipher.NewCBCDecrypter(block, iv).CryptBlocks(con.Key, con.Key)
      con.Key = unpad(con.Key)
      cons = append(cons, con)
   }
   return cons, nil
}
