package widevine
// github.com/89z

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
   ID []byte
   PrivateKey []byte
   RawKeyID string
   RawPSSH string
}

func (c Client) KeyID() ([]byte, error) {
   if c.RawKeyID != "" {
      c.RawKeyID = strings.ReplaceAll(c.RawKeyID, "-", "")
      return hex.DecodeString(c.RawKeyID)
   }
   _, after, ok := strings.Cut(c.RawPSSH, "data:text/plain;base64,")
   if ok {
      c.RawPSSH = after
   }
   pssh, err := base64.StdEncoding.DecodeString(c.RawPSSH)
   if err != nil {
      return nil, err
   }
   widevineCencHeader, err := protobuf.Unmarshal(pssh[32:])
   if err != nil {
      return nil, err
   }
   return widevineCencHeader.GetBytes(2)
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
