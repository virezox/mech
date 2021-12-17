package pandora

import (
   "encoding/hex"
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "strings"
   "time"
)

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

// For some reason the UserAuthToken being returned by this doesnt actually
// work.
func newUserLogin(username, password string) (*userLogin, error) {
   syncTime := time.Now().Unix()
   partnerAuthToken := ""
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
   `, syncTime, partnerAuthToken)
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
