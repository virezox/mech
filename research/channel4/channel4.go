package channel4

import (
   "encoding/base64"
   "strings"
)

func createPSSH(kid string) string {
   buf := []byte{
      0,    0,    0,    '2', 'p',   's',  's', 'h',
      0,    0,    0,    0,
      0xed, 0xef, 0x8b, 0xa9, 0x79, 0xd6, 0x4a, 0xce,
      0xa3, 0xc8, 0x27, 0xdc, 0xd5, 0x1d, 0x21, 0xed,
      0,    0,    0,    0x12, 0x12, 0x10,
   }
   kid = strings.ReplaceAll(kid, "-", "")
   buf = append(buf, kid...)
   return base64.StdEncoding.EncodeToString(buf)
}
