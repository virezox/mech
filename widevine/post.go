package widevine

import (
   "bytes"
   "crypto"
   "crypto/aes"
   "crypto/cipher"
   "crypto/rsa"
   "crypto/sha1"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/protobuf"
   "github.com/chmike/cmac-go"
   "io"
)

type Poster interface {
   Request_URL() string
   Request_Header() http.Header
   Request_Body([]byte) ([]byte, error)
   Response_Body([]byte) ([]byte, error)
}

func (self Module) Post(post Poster) (Containers, error) {
   signed_request, err := self.signed_request()
   if err != nil {
      return nil, err
   }
   body, err := post.Request_Body(signed_request)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", post.Request_URL(), bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   if h := post.Request_Header(); h != nil {
      req.Header = h
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   body, err = post.Response_Body(body)
   if err != nil {
      return nil, err
   }
   return self.signed_response(body)
}

func (self Module) signed_response(response []byte) (Containers, error) {
   // key
   signed_response, err := protobuf.Unmarshal(response)
   if err != nil {
      return nil, err
   }
   session_key, err := signed_response.Get_Bytes(4)
   if err != nil {
      return nil, err
   }
   key, err := rsa.DecryptOAEP(
      sha1.New(), nil, self.private_key, session_key, nil,
   )
   if err != nil {
      return nil, err
   }
   // message
   var b []byte
   b = append(b, 1)
   b = append(b, "ENCRYPTION"...)
   b = append(b, 0)
   b = append(b, self.license_request...)
   b = append(b, 0, 0, 0, 0x80)
   // CMAC
   mac, err := cmac.New(aes.NewCipher, key)
   if err != nil {
      return nil, err
   }
   mac.Write(b)
   block, err := aes.NewCipher(mac.Sum(nil))
   if err != nil {
      return nil, err
   }
   var cons Containers
   // .Msg.Key
   for _, message := range signed_response.Get(2).Get_Messages(3) {
      var con Container
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

func (self Module) signed_request() ([]byte, error) {
   digest := sha1.Sum(self.license_request)
   signature, err := rsa.SignPSS(
      no_operation{},
      self.private_key,
      crypto.SHA1,
      digest[:],
      &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash},
   )
   if err != nil {
      return nil, err
   }
   signed_request := protobuf.Message{
      2: protobuf.Bytes(self.license_request),
      3: protobuf.Bytes(signature),
   }
   return signed_request.Marshal(), nil
}
