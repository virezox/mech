package pandora

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

type playbackInfo struct {
   Result struct {
      AudioUrlMap struct {
         HighQuality struct {
            AudioURL string
         }
      }
   }
}

func (u userLogin) playbackInfo() (*playbackInfo, error) {
   // this is one i created that works:
   userAuthToken := "VIfvY3SVp7VlmLbuzm+S6er6mEYuah54TTkfGLOnrRpK0ay+vEpcXvcQ=="
   deviceCode := ""
   dec := fmt.Sprintf(`
  {
   "userAuthToken": "%v",
   "includeAudioToken": true,
   "pandoraId": "TR:1168891",
   "deviceCode": "%v",
   "syncTime": 2222222222
}
   `, userAuthToken, deviceCode)
   enc, err := encrypt([]byte(dec))
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
   // this can be empty, but it must be included:
   val["auth_token"] = ""
   val["method"] = "onDemand.getAudioPlaybackInfo"
   val["partner_id"] = "42"
   // this can be empty, but it must be included:
   val["user_id"] = ""
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(buf)
   info := new(playbackInfo)
   if err := json.NewDecoder(res.Body).Decode(info); err != nil {
      return nil, err
   }
   return info, nil
}

func newUserLogin(partnerAuthToken string, body []byte) (*userLogin, error) {
   enc, err := encrypt(body)
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
   val["auth_token"] = partnerAuthToken
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
