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

const origin = "http://android-tuner.pandora.com"

var (
   LogLevel mech.LogLevel
   key = []byte("6#26FRL$ZWD")
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

type playbackInfoRequest struct {
   // this can be empty, but must be included:
   DeviceCode string `json:"deviceCode"`
   IncludeAudioToken bool `json:"includeAudioToken"`
   PandoraID string `json:"pandoraId"`
   SyncTime int64 `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}

type userLoginRequest struct {
   LoginType string `json:"loginType"`
   PartnerAuthToken string `json:"partnerAuthToken"`
   Password string `json:"password"`
   SyncTime int64 `json:"syncTime"`
   Username string `json:"username"`
}

func deviceID() (*userLogin, error) {
   // how to get deviceId?
   deviceID := "7db1bef0-1ea2-4ba5-8c0a-079274c81b75"
   partnerAuthToken := "VAue2IpSC8HwJuTW3NK4TkdT3Irv9FrXZF"
   body := fmt.Sprintf(`
   {
      "deviceId": "%v",
      "loginType": "deviceId",
      "partnerAuthToken": "%v",
      "syncTime": 2222222222
   }
   `, deviceID, partnerAuthToken)
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

