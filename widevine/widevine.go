package widevine
// github.com/89z

import (
   "encoding/base64"
   "encoding/hex"
   "github.com/89z/format/protobuf"
   "strings"
)

func KeyID(rawPSSH string) ([]byte, error) {
   _, after, ok := strings.Cut(rawPSSH, "data:text/plain;base64,")
   if ok {
      rawPSSH = after
   }
   pssh, err := base64.StdEncoding.DecodeString(rawPSSH)
   if err != nil {
      return nil, err
   }
   widevineCencHeader, err := protobuf.Unmarshal(pssh[32:])
   if err != nil {
      return nil, err
   }
   // key_id
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

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
}
