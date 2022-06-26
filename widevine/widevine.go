package widevine

import (
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "encoding/hex"
   "github.com/89z/format/protobuf"
   "github.com/chmike/cmac-go"
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

type Content struct {
   Key []byte
   Type uint64
}

func (c Content) String() string {
   return hex.EncodeToString(c.Key)
}

type Contents []Content

func (c Contents) Content() *Content {
   for _, con := range c {
      if con.Type == 2 {
         return &con
      }
   }
   return nil
}

type Module struct {
   license_request []byte
   private_key *rsa.PrivateKey
}

func (m Module) Marshal() ([]byte, error) {
   digest := sha1.Sum(m.license_request)
   signature, err := rsa.SignPSS(
      no_operation{},
      m.private_key,
      crypto.SHA1,
      digest[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   signed_request := protobuf.Message{
      2: protobuf.Bytes(m.license_request),
      3: protobuf.Bytes(signature),
   }
   return signed_request.Marshal(), nil
}

func (m Module) Unmarshal(response []byte) (Contents, error) {
   // key
   signed_response, err := protobuf.Unmarshal(response)
   if err != nil {
      return nil, err
   }
   session_key, err := signed_response.Get_Bytes(4)
   if err != nil {
      return nil, err
   }
   key, err := rsa.DecryptOAEP(sha1.New(), nil, m.private_key, session_key, nil)
   if err != nil {
      return nil, err
   }
   // message
   var mes []byte
   mes = append(mes, 1)
   mes = append(mes, "ENCRYPTION"...)
   mes = append(mes, 0)
   mes = append(mes, m.license_request...)
   mes = append(mes, 0, 0, 0, 0x80)
   // CMAC
   mac, err := cmac.New(aes.NewCipher, key)
   if err != nil {
      return nil, err
   }
   mac.Write(mes)
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

type no_operation struct{}

func (no_operation) Read(buf []byte) (int, error) {
   return len(buf), nil
}
