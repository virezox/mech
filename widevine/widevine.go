package widevine

import (
   "encoding/base64"
   "encoding/hex"
   "github.com/89z/format/protobuf"
   "strings"
)

func (c Content) String() string {
   return hex.EncodeToString(c.Key)
}

type Content struct {
   Key []byte
   Type uint64
}

type Client struct {
   Id []byte
   PrivateKey []byte
   RawKeyId string
   RawPssh string
}

func (c Client) KeyId() ([]byte, error) {
   if c.RawKeyId != "" {
      c.RawKeyId = strings.ReplaceAll(c.RawKeyId, "-", "")
      return hex.DecodeString(c.RawKeyId)
   }
   _, after, ok := strings.Cut(c.RawPssh, "data:text/plain;base64,")
   if ok {
      c.RawPssh = after
   }
   pssh, err := base64.StdEncoding.DecodeString(c.RawPssh)
   if err != nil {
      return nil, err
   }
   cencHeader := make(protobuf.Message)
   if err := cencHeader.UnmarshalBinary(pssh[32:]); err != nil {
      return nil, err
   }
   return cencHeader.GetBytes(2)
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

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
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
