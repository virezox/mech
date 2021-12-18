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

func newUser() (*userLogin, error) {
   partnerAuthToken := "VAa8OzW17TN/KS1tKZKzYjxY63lzaFGvi/"
   body := fmt.Sprintf(`
   {
   "loginType": "user",
   "partnerAuthToken": "%v",
   "password": "%v",
   "syncTime": 2222222222,
   "username": "srpen6@gmail.com"
}
   `, partnerAuthToken, password)
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
   req.Header.Set("User-Agent", "Pandora/2110.1 Android/7.0 generic_x86")
   val := make(mech.Values)
   val["auth_token"] = ""
   val["method"] = "auth.userLogin"
   val["partner_id"] = "42"
   req.URL.RawQuery = val.Encode()
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

var (
   LogLevel mech.LogLevel
   key = []byte("6#26FRL$ZWD")
)

type userLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
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
   req.Header.Set("User-Agent", "Pandora/2110.1 Android/7.0 generic_x86")
   req.URL.RawQuery = "method=auth.partnerLogin"
   LogLevel.Dump(req)
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
