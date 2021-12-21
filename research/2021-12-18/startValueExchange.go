package main

import (
   "encoding/hex"
   "fmt"
   "github.com/89z/mech"
   "golang.org/x/crypto/blowfish"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)

var (
   LogLevel mech.LogLevel
   key = []byte("6#26FRL$ZWD")
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

func main() {
   if len(os.Args) != 2 {
      fmt.Println("startValueExchange [userAuthToken]")
      return 
   }
   userAuthToken := os.Args[1]
   body := fmt.Sprintf(`
   {
   "userAuthToken": "%v",
   "lineId": "39569",
   "engagementCompleted": true,
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
   "adServerCorrelationId": "3517222595310218723",
   "includeFlexParams": true,
   "syncTime": 1639854438,
   "useRewardAfterCreation": true,
   "creativeId": "54012",
   "rewardProperties": {
   "playableSourceId": "TR:1168891",
   "playableSource": "TR",
   "campaignId": "P1429568",
   "lineId": "39569",
   "playableTarget": "TR",
   "advertiserId": "Haymaker",
   "playableTargetId": "TR:1168891",
   "placementId": "SLOPA",
   "siteId": "android",
   "learnMoreUrl": "https://ad.doubleclick.net/ddm/trackclk/N6249.127425.PANDORA/B25211457.291538383;dc_trk_aid=508394281;dc_trk_cid=126843437;dc_lat=;dc_rdid=;tag_for_child_directed_treatment=;tfua=;ltd=",
   "creativeId": "54012"
   },
   "offerName": "premium_access"
   }
   `, userAuthToken)
   enc, err := encrypt([]byte(body))
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST", "http://android-tuner.pandora.com/services/json/",
      strings.NewReader(hex.EncodeToString(enc)),
   )
   if err != nil {
      panic(err)
   }
   val := make(mech.Values)
   // this can be empty, but must be included:
   val["auth_token"] = ""
   val["method"] = "user.startValueExchange"
   val["partner_id"] = "42"
   // this can be empty, but it must be included:
   val["user_id"] = ""
   req.URL.RawQuery = val.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}
