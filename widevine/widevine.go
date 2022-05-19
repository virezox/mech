package widevine

import (
   "encoding/hex"
   "github.com/89z/format"
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

type nopSource struct{}

func (nopSource) Read(buf []byte) (int, error) {
   return len(buf), nil
}
