package widevine

import (
   "crypto/x509"
   "encoding/base64"
   "encoding/hex"
   "encoding/pem"
   "github.com/89z/format/protobuf"
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
