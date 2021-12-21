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

type partnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}

func newPartnerLogin() (*partnerLogin, error) {
   body := strings.NewReader(`
   {
  "returnUpdatePromptVersions": true,
  "deviceModel": "android-generic_x86",
  "version": "5",
  "advertisingTrackingEnabled": "YES",
  "deviceProperties": {
    "deviceCategory": "android",
    "w": "1080",
    "model": "android-generic_x86",
    "applicationVersionCode": "21101001",
    "carrierName": "Android",
    "isFromAmazon": "false",
    "h": "1794",
    "code": "android-generic_x86",
    "applicationVersion": "2110.1",
    "systemVersion": "7.0",
    "fordInfo": "{HMIStatus=NONE}"
  },
  "password": "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
  "returnDeviceType": true,
  "deviceTrackingIds": [
    "471b41d1-25f4-4a2f-881a-80c8b693da2c",
    "c4e64ee06038bfca",
    "fa1fa21f-1458-4b49-bfca-87ac84005f15"
  ],
  "includeExtraParams": true,
  "includeUrls": true,
  "username": "android"
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

func newUserLogin() (*userLogin, error) {
   partnerAuthToken := "VAVKoVBC3zhbKAD/zvE1daYyALJB7VXt3C"
   body := fmt.Sprintf(`
   {
      "deviceId": "7db1bef0-1ea2-4ba5-8c0a-079274c81b75",
      "loginType": "deviceId",
      "partnerAuthToken": "%v",
      "syncTime": 2222222222
   }
   `, partnerAuthToken)
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
