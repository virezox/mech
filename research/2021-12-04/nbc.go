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

func media(guid int) (*http.Response, error) {
   req, err := http.NewRequest(
      "GET",
      platform + "/s/NnzsPC/media/guid/2410887629/" + strconv.Itoa(guid),
      nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      //"format": {"SMIL"}, // can kill
      //"manifest": {"m3u"}, // maybe can kill?
      "mbr": {"true"}, // can kill
   }.Encode()
   mech.Dump(req)
   return new(http.Transport).RoundTrip(req)
}

func newAccessVOD(guid int) (*accessVOD, error) {
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
      "POST", origin + "/access/vod/nbcuniversal/" + strconv.Itoa(guid), buf,
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


