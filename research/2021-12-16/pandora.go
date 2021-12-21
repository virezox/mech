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
   "time"
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

func newPlaybackInfo() (*playbackInfo, error) {
   userAuthToken := "VIB/H5cN0QRf34p6atwS0N7uCtcUp4wRL00iVw4aG6kOUYU9RCy3vegA=="
   dec := fmt.Sprintf(`
  {
   "userAuthToken": "%v",
   "includeAudioToken": true,
   "pandoraId": "TR:1168891",
   "deviceCode": "",
   "syncTime": 2222222222
}
   `, userAuthToken)
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

type userLogin struct {
   Result struct {
      UserID string
      UserAuthToken string
   }
}

// For some reason the UserAuthToken being returned by this doesnt actually
// work.
func (p partnerLogin) userLogin(username, password string) (*userLogin, error) {
   dec := fmt.Sprintf(`
{
 "returnGenreStations": false,
 "includeShuffleInsteadOfQuickMix": true,
 "includeFlexParams": true,
 "includeAccountMessage": true,
 "stationArtSize": "W500H500",
 "includeFacebook": true,
 "includeSlopaAdUrl": true,
 "includeListeningHours": true,
 "includePlaylistAttributes": true,
 "includeStationArtUrl": true,
 "includeSkipAttributes": true,
 "advertisingTrackingEnabled": "YES",
 "deviceTrackingIds": [
  "72d81533-15bf-4c1c-b9b6-04c39e980db0",
  "c4e64ee06038bfca",
  "fa1fa21f-1458-4b49-bfca-87ac84005f15"
 ],
 "stationListChecksum": "648bda20223d4bfa22fe80190f54a666",
 "deviceId": "72d81533-15bf-4c1c-b9b6-04c39e980db0",
 "includeStationSeeds": true,
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
 "includeDemographics": true,
 "includeGoogleplay": true,
 "includeStatsCollectorConfig": true,
 "returnAllStations": true,
 "locale": "en_US",
 "returnCollectTrackLifetimeStats": true,
 "includeUserWebname": true,
 "includeDailySkipLimit": true,
 "xplatformAdCapable": true,
 "premiumCapable": true,
 "complimentarySponsorSupported": true,
 "includeStationDescription": true,
 "includeAdAttributes": true,
 "includePandoraOneInfo": true,
 "returnUserstate": true,
 "includeSlopaNoAvailsAdUrl": true,
 "includeGenreCategoryAdUrl": true,
 "includeSubscriptionExpiration": true,
 "returnHasUsedTrial": true,
 "shuffleIconVersion": 2,
 "returnStationList": true,
 "includeRewardedAdUrl": true,
 "returnIsSubscriber": true,
 "includeExtraParams": true,
 "loginType": "deviceId",
 "includeTwitter": true,
 "includeAdvertiserAttributes": true,
 "includeABTesting": true,
 "includeStationExpirationTime": true,
 "includeSkipDelay": true,
 "syncTime": %v,
 "partnerAuthToken": %q,
}
   `, time.Now().Unix(), p.Result.PartnerAuthToken)
   enc, err := encrypt([]byte(dec))
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
   req.Header.Set("User-Agent", "Pandora/2110.1 Android/7.0 generic_x86")
   val := make(mech.Values)
   val["auth_token"] = p.Result.PartnerAuthToken
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
