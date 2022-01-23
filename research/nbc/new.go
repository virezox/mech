package main

import (
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "github.com/89z/format"
   "io"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

var (
   LogLevel format.LogLevel
   secretKey = []byte("2b84a073ede61c766e4c0b3f1e656f7f")
)

func generateHash(text string, key []byte) string {
   mac := hmac.New(sha256.New, key)
   io.WriteString(mac, text)
   sum := mac.Sum(nil)
   return hex.EncodeToString(sum)
}

var body = strings.NewReader(`
{
  "device": "android",
  "auth": false,
  "adobeMvpdId": "",
  "deviceId": "e3a71011a55b81b9",
  "device_type": "google_advertising_id",
  "nw": "169843",
  "mParticleId": "-7855216639727670072",
  "externalAdvertiserId": "NBC_VOD_9000199368",
  "did": "e3a71011a55b81b9",
  "uuid": "76004574-fc3f-4683-9430-df6598dda1fb",
  "appv": "NBC_7.28.1",
  "buildv": "7.28.1",
  "am_appv": "7.28.1",
  "am_buildv": "2000002882",
  "player_height": "1080",
  "player_width": "1920",
  "sdkv": "android_4.10.71.2",
  "playerv": "exoplayer_2.11.8",
  "bundleId": "com.nbcuni.nbc",
  "us_privacy_string": "",
  "mpx": {
    "accountId": "2304985974"
  },
  "tracking": {
    "deviceGroup": "PHN",
    "platform": "MBL",
    "appId": "PAD3C6E72-ED61-417F-A865-3AB63FDB6197",
    "androidId": "76004574-fc3f-4683-9430-df6598dda1fb",
    "comscore_device": "Android_SDK_built_for_x86",
    "googleAdId": "76004574-fc3f-4683-9430-df6598dda1fb"
  },
  "prefetch": false
}
`)

func main() {
   var req http.Request
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"application/access-v1+json"}
   req.Header["Accept-Encoding"] = []string{"gzip"}
   req.Header["App-Session-Id"] = []string{"dc5dac8c-ca7f-4ef6-88f8-0dc529bacc46"}
   unix := strconv.FormatInt(time.Now().UnixMilli(), 10)
   var auth strings.Builder
   auth.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   auth.WriteString(",hash=")
   auth.WriteString(generateHash(unix, secretKey))
   auth.WriteString(",time=")
   auth.WriteString(unix)
   req.Header["Authorization"] = []string{auth.String()}
   req.Header["Cache-Control"] = []string{"no-cache"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Length"] = []string{"815"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Host"] = []string{"access-cloudpath.media.nbcuni.com"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 7.0; Android SDK built for x86 Build/NYC; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 Mobile Safari/537.36"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "access-cloudpath.media.nbcuni.com"
   //pass
   req.URL.Path = "/access/vod/nbcuniversal/9000199368"
   //fail
   //req.URL.Path = "/access/vod/nbcuniversal/9000199367"
   val := make(url.Values)
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "http"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   file, err := os.Create("res.json")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   file.ReadFrom(res.Body)
}

