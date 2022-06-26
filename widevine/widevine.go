package widevine

import (
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/base64"
   "encoding/hex"
   "encoding/pem"
   "github.com/89z/format/protobuf"
   "github.com/chmike/cmac-go"
   "strings"
)

type Client struct {
   ID []byte
   Private_Key []byte
   Raw string
}

func (c Client) module(key_id []byte) (*Module, error) {
   block, _ := pem.Decode(c.Private_Key)
   var (
      err error
      mod Module
   )
   mod.private_key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   mod.license_request = protobuf.Message{
      1: protobuf.Bytes(c.ID),
      2: protobuf.Message{ // ContentId
         1: protobuf.Message{ // CencId
            1: protobuf.Message{ // Pssh
               2: protobuf.Bytes(key_id),
            },
         },
      },
   }.Marshal()
   return &mod, nil
}

func (c Client) Key_ID() (*Module, error) {
   c.Raw = strings.ReplaceAll(c.Raw, "-", "")
   key_id, err := hex.DecodeString(c.Raw)
   if err != nil {
      return nil, err
   }
   return c.module(key_id)
}

func (c Client) PSSH() (*Module, error) {
   _, after, ok := strings.Cut(c.Raw, "data:text/plain;base64,")
   if ok {
      c.Raw = after
   }
   pssh, err := base64.StdEncoding.DecodeString(c.Raw)
   if err != nil {
      return nil, err
   }
   cenc_header, err := protobuf.Unmarshal(pssh[32:])
   if err != nil {
      return nil, err
   }
   key_id, err := cenc_header.Get_Bytes(2)
   if err != nil {
      return nil, err
   }
   return c.module(key_id)
}
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
