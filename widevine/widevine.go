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
   ID []byte
   Private_Key []byte
   Raw_Key_ID string
   Raw_PSSH string
}

func (c Client) Key_ID() ([]byte, error) {
   if c.Raw_Key_ID != "" {
      c.Raw_Key_ID = strings.ReplaceAll(c.Raw_Key_ID, "-", "")
      return hex.DecodeString(c.Raw_Key_ID)
   }
   _, after, ok := strings.Cut(c.Raw_PSSH, "data:text/plain;base64,")
   if ok {
      c.Raw_PSSH = after
   }
   pssh, err := base64.StdEncoding.DecodeString(c.Raw_PSSH)
   if err != nil {
      return nil, err
   }
   cenc_header, err := protobuf.Unmarshal(pssh[32:])
   if err != nil {
      return nil, err
   }
   return cenc_header.Get_Bytes(2)
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
