package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/format"
   "io"
   "net/http"
   "strconv"
   "strings"
   "time"
)

const mpxAccountID = 2410887629

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

type AccessVOD struct {
   ManifestPath string // this is only valid for one minute
}

func NewAccessVOD(guid uint64) (*AccessVOD, error) {
   var body vodRequest
   body.Device = "android"
   body.DeviceID = "android"
   body.ExternalAdvertiserID = "NBC"
   body.Mpx.AccountID = mpxAccountID
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   addr := []byte("http://access-cloudpath.media.nbcuni.com")
   addr = append(addr, "/access/vod/nbcuniversal/"...)
   addr = strconv.AppendUint(addr, guid, 10)
   req, err := http.NewRequest("POST", string(addr), buf)
   if err != nil {
      return nil, err
   }
   unix := strconv.FormatInt(time.Now().UnixMilli(), 10)
   var auth strings.Builder
   auth.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   auth.WriteString(",time=")
   auth.WriteString(unix)
   auth.WriteString(",hash=")
   auth.WriteString(generateHash(unix, secretKey))
   req.Header = http.Header{
      "Authorization": {auth.String()},
      "Content-Type": {"application/json"},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vod := new(AccessVOD)
   if err := json.NewDecoder(res.Body).Decode(vod); err != nil {
      return nil, err
   }
   return vod, nil
}

type vodRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID int `json:"accountId"`
   } `json:"mpx"`
}
