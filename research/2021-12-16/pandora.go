package pandora

import (
   "encoding/json"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "strings"
)

type partnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}

func newPartnerLogin() (*partnerLogin, error) {
   body := `{"username":"android","password":"AC7IBG09A3DTSYM4R41UJWL07VLN8JI7"}`
   req, err := http.NewRequest(
      "POST", "http://android-tuner.pandora.com/services/json/",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "method=auth.partnerLogin"
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   login := new(partnerLogin)
   if err := json.NewDecoder(res.Body).Decode(login); err != nil {
      return nil, err
   }
   return login, nil
}

var (
   keyDecrypt = []byte("6#26FRL$ZWD")
   keyEncrypt = []byte("R=U!LH$O2B#")
)

func decrypt(src []byte) ([]byte, error) {
   dst := make([]byte, len(src))
   blow, err := blowfish.NewCipher(keyDecrypt)
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
   pad := make(
      []byte, blowfish.BlockSize-len(src)%blowfish.BlockSize,
   )
   src = append(src, pad...)
   dst := make([]byte, len(src))
   blow, err := blowfish.NewCipher(keyEncrypt)
   if err != nil {
      return nil, err
   }
   for low := 0; low < len(src); low += blow.BlockSize() {
      blow.Encrypt(dst[low:], src[low:])
   }
   return dst, nil
}
