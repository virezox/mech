package widevine

import (
   "bytes"
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "crypto/x509"
   "encoding/hex"
   "encoding/pem"
   "errors"
   "github.com/89z/format"
   "github.com/89z/format/protobuf"
   "github.com/chmike/cmac-go"
   "io"
   "net/http"
)

var LogLevel format.LogLevel

func unpad(buf []byte) []byte {
   if len(buf) >= 1 {
      pad := buf[len(buf)-1]
      if len(buf) >= int(pad) {
         buf = buf[:len(buf)-int(pad)]
      }
   }
   return buf
}

type Container struct {
   Key []byte
   Type uint64
}

func (c Container) String() string {
   return hex.EncodeToString(c.Key)
}

type Containers []Container

func (c Containers) Content() *Container {
   for _, con := range c {
      if con.Type == 2 {
         return &con
      }
   }
   return nil
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

func (m Module) Keys(body io.Reader) (Containers, error) {
   // key
   signedResponse, err := protobuf.Decode(body)
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
      var con Container
      con.Key = unpad(key)
      con.Type = uint64(typ)
      cons = append(cons, con)
   }
   return cons, nil
}

func (m Module) Post(addr string, head http.Header) (Containers, error) {
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
      2: protobuf.Bytes(m.licenseRequest),
      3: protobuf.Bytes(signature),
   }
   req, err := http.NewRequest(
      "POST", addr, bytes.NewReader(signedRequest.Marshal()),
   )
   if err != nil {
      return nil, err
   }
   req.Header = head
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   return m.Keys(res.Body)
}

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
}
