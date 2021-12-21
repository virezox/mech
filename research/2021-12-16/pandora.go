package pandora

import (
   "encoding/hex"
   "encoding/json"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
   "time"
)

var (
   _ = hex.EncodeToString
   _ = time.Sleep
)

type values map[string]string

func (v values) encode() string {
   vals := make(url.Values)
   for key, val := range v {
      vals.Set(key, val)
   }
   return vals.Encode()
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

type partnerLogin struct {
   Result struct {
      PartnerAuthToken string
   }
}

func newPartnerLogin() (*partnerLogin, error) {
   body := `{"username":"android","password":"AC7IBG09A3DTSYM4R41UJWL07VLN8JI7"}`
   req, err := http.NewRequest(
      "POST", origin + "/services/json/", strings.NewReader(body),
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

type playbackInfo struct {
   Result struct {
      AudioUrlMap struct {
         HighQuality struct {
            AudioURL string
         }
      }
   }
}

type playbackInfoRequest struct {
   // this can be empty, but must be included:
   DeviceCode string `json:"deviceCode"`
   IncludeAudioToken bool `json:"includeAudioToken"`
   PandoraID string `json:"pandoraId"`
   SyncTime int64 `json:"syncTime"`
   UserAuthToken string `json:"userAuthToken"`
}

func (u userLogin) playbackInfo() (*playbackInfo, error) {
   rInfo := playbackInfoRequest{
      IncludeAudioToken: true,
      SyncTime: 2222222222,
      PandoraID: "TR:1168891",
      // this is failing:
      UserAuthToken: u.Result.UserAuthToken,
   }
   dec, err := json.Marshal(rInfo)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(dec)
   enc, err := encrypt(dec)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
   )
   if err != nil {
      return nil, err
   }
   // method=auth.partnerLogin
   req.URL.RawQuery = values{
      // this can be empty, but it must be included:
      "auth_token": "",
      "method": "onDemand.getAudioPlaybackInfo",
      "partner_id": "42",
      // this can be empty, but it must be included:
      "user_id": "",
   }.encode()
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
