package pandora

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
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

type userLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

func (p partnerLogin) userLogin(username, password string) (*userLogin, error) {
   body := fmt.Sprintf(`
   {
      "loginType": "user",
      "partnerAuthToken": "%v",
      "password": "%v",
      "syncTime": 2222222222,
      "username": "srpen6@gmail.com"
   }
   `, p.Result.PartnerAuthToken, password)
   enc, err := encrypt([]byte(body))
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST",
      "http://android-tuner.pandora.com/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
   )
   if err != nil {
      return nil, err
   }
   val := make(mech.Values)
   // this can be empty, but must be included:
   val["auth_token"] = ""
   val["method"] = "auth.userLogin"
   val["partner_id"] = "42"
   req.URL.RawQuery = val.Encode()
   logLevel.Dump(req)
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
