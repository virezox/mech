package pandora

import (
   "encoding/json"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "strings"
)

var (
   key = []byte("6#26FRL$ZWD")
   logLevel mech.LogLevel
)

func decrypt(src []byte) ([]byte, error) {
   dst := make([]byte, len(src))
   blow, err := blowfish.NewCipher(key)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(src); low += blowfish.BlockSize {
      blow.Decrypt(dst[low:], src[low:])
   }
   pad := dst[len(dst)-1]
   return dst[:len(dst)-int(pad)], nil
}

func encrypt(src []byte) ([]byte, error) {
   for len(src) % blowfish.BlockSize >= 1 {
      src = append(src, 0)
   }
   dst := make([]byte, len(src))
   blow, err := blowfish.NewCipher(key)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(src); low += blow.BlockSize() {
      blow.Encrypt(dst[low:], src[low:])
   }
   return dst, nil
}

type partnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}

func newPartnerLogin() (*partnerLogin, error) {
   body := strings.NewReader(`
   {
      "deviceModel": "android-generic_x86",
      "password": "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
      "username": "android",
      "version": "5"
   }
   `)
   req, err := http.NewRequest(
      "POST", "http://android-tuner.pandora.com/services/json/", body,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "method=auth.partnerLogin"
   logLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   part := new(partnerLogin)
   if err := json.NewDecoder(res.Body).Decode(part); err != nil {
      return nil, err
   }
   return part, nil
}
