package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const platform = "http://link.theplatform.com"

func media() (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", platform + "/s/NnzsPC/media/guid/2410887629/9000194212", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "format": {"SMIL"}, // can kill
      "mbr": {"true"}, // can kill
      //"manifest": {"m3u"}, // maybe can kill?
   }.Encode()
   return new(http.Transport).RoundTrip(req)
}

const (
   mpxAccountID = 2304985974
   origin = "http://access-cloudpath.media.nbcuni.com"
)

var secretKey = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

func generateHash(text, key []byte) string {
   mac := hmac.New(sha256.New, key)
   mac.Write(text)
   sum := mac.Sum(nil)
   return hex.EncodeToString(sum)
}

func unixMilli() []byte {
   unix := time.Now().UnixMilli()
   return strconv.AppendInt(nil, unix, 10)
}

type accessVOD struct {
   // this is only valid for one minute
   ManifestPath string
}

func newAccessVOD(id int) (*accessVOD, error) {
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
   req, err := http.NewRequest(
      "POST", origin + "/access/vod/nbcuniversal/" + strconv.Itoa(id), buf,
   )
   if err != nil {
      return nil, err
   }
   unix := unixMilli()
   var auth strings.Builder
   auth.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   auth.WriteString(",hash=")
   auth.WriteString(generateHash(unix, secretKey))
   auth.WriteString(",time=")
   auth.Write(unix)
   // If you add "User-Agent: Mozilla/5", more information will be returned.
   req.Header = http.Header{
      "authorization": {auth.String()},
      "content-type": {"application/json"},
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vod := new(accessVOD)
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
