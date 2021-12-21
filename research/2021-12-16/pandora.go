package pandora

import (
   "encoding/hex"
   "encoding/json"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "net/url"
   "strings"
)

func encrypt(src []byte) ([]byte, error) {
   for len(src) % blowfish.BlockSize >= 1 {
      src = append(src, 0)
   }
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

func (p partnerLogin) userLogin() (*userLogin, error) {
   buf, err := encrypt(userLoginDec)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "http://android-tuner.pandora.com/services/json/",
      //strings.NewReader(userLoginEnc),
      strings.NewReader(hex.EncodeToString(buf)),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery =
      "partner_id=42&method=auth.userLogin&auth_token=" +
      url.QueryEscape(p.Result.PartnerAuthToken)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   user := new(userLogin)
   if err := json.NewDecoder(res.Body).Decode(user); err != nil {
      return nil, err
   }
   return user, nil
}

var LogLevel mech.LogLevel

type partnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}

type userLogin struct {
   Result struct {
      UserID string
   }
}

var (
   keyDecrypt = []byte("6#26FRL$ZWD")
   keyEncrypt = []byte("R=U!LH$O2B#")
)

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
   LogLevel.Dump(req)
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
