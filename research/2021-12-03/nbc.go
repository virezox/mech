package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "strings"
   "time"
)

const origin = "http://access-cloudpath.media.nbcuni.com"

type vodRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID int `json:"accountId"`
   } `json:"mpx"`
}

var key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

func unixBytes() []byte {
   unix := time.Now().UnixMilli()
   return strconv.AppendInt(nil, unix, 10)
}

func vod() (*http.Response, error) {
   var body vodRequest
   body.Device = "android"
   body.DeviceID = "android"
   body.ExternalAdvertiserID = "NBC"
   body.Mpx.AccountID = 2304985974
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", origin + "/access/vod/nbcuniversal/9000194212", buf,
   )
   if err != nil {
      return nil, err
   }
   unix := unixBytes()
   var auth strings.Builder
   auth.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   auth.WriteString(",hash=")
   auth.WriteString(generateHash(unix, key))
   auth.WriteString(",time=")
   auth.Write(unix)
   req.Header = http.Header{
      "authorization": {auth.String()},
      "content-type": {"application/json"},
      "user-agent": {"Mozilla/5"},
   }
   mech.Verbose = true
   mech.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

func generateHash(text, key []byte) string {
   mac := hmac.New(sha256.New, key)
   mac.Write(text)
   sum := mac.Sum(nil)
   return hex.EncodeToString(sum)
}
