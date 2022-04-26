package channel4

import (
   "encoding/base64"
   "encoding/hex"
   "strings"
)

func createPSSH(kid string) (string, error) {
   a := []byte{0, 0, 0, '2', 'p', 's', 's', 'h', 0, 0, 0, 0}
   b, err := hex.DecodeString("edef8ba979d64acea3c827dcd51d21ed")
   if err != nil {
      return "", err
   }
   a = append(a, b...)
   a = append(a, 0, 0, 0, 0x12, 0x12, 0x10)
   kid = strings.ReplaceAll(kid, "-", "")
   c, err := hex.DecodeString(kid)
   if err != nil {
      return "", err
   }
   a = append(a, c...)
   return base64.StdEncoding.EncodeToString(a), nil
}
