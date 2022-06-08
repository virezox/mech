package widevine
// github.com/89z

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

func (m Module) Unmarshal(response []byte) (Containers, error) {
   // key
   signedResponse, err := protobuf.Unmarshal(response)
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
   var cons Containers
   // .Msg.Key
   for _, message := range signedResponse.Get(2).GetMessages(3) {
      var con Container
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
      1: protobuf.Bytes{Raw: clientID},
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes{Raw: kID},
            },
         },
      },
   }.Marshal()
   return &mod, nil
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
   return signedRequest.Marshal(), nil
}
