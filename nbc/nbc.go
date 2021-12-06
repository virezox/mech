package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "strconv"
   "strings"
   "time"
)

const (
   Origin = "http://access-cloudpath.media.nbcuni.com"
   queryVideo = "73014253e5761c29fc76b950e7d4d181c942fa401b3378af4bac366f6611601e"
)

const accountID = 2410887629

var secretKey = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

// nbc.com/la-brea/video/pilot/9000194212
func Valid(guid string) bool {
   return len(guid) == 10
}

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

type AccessVOD struct {
   // this is only valid for one minute
   ManifestPath string
}

func NewAccessVOD(guid int) (*AccessVOD, error) {
   var body vodRequest
   body.Device = "android"
   body.DeviceID = "android"
   body.ExternalAdvertiserID = "NBC"
   body.Mpx.AccountID = accountID
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", Origin + "/access/vod/nbcuniversal/" + strconv.Itoa(guid), buf,
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
   req.Header = http.Header{
      "Authorization": {auth.String()},
      "Content-Type": {"application/json"},
   }
   mech.Dump(req)
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

func (a AccessVOD) Manifest() ([]m3u.Format, error) {
   req, err := http.NewRequest("GET", a.ManifestPath, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return m3u.Decode(res.Body, "")
}

type Video struct {
   Data struct {
      BonanzaPage struct {
         Analytics struct {
            ConvivaAssetName string
         }
      }
   }
}

func NewVideo(guid int) (*Video, error) {
   var rVid videoRequest
   rVid.Extensions.PersistedQuery.Sha256Hash = queryVideo
   rVid.Variables.App = "nbc"
   rVid.Variables.Name = strconv.Itoa(guid)
   rVid.Variables.Platform = "android"
   rVid.Variables.Type = "VIDEO"
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(rVid)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://friendship.nbc.co/v2/graphql", buf,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   mech.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

func (v Video) Name() string {
   return v.Data.BonanzaPage.Analytics.ConvivaAssetName
}

type videoRequest struct {
   Extensions struct {
      PersistedQuery struct {
         Sha256Hash string `json:"sha256Hash"`
      } `json:"persistedQuery"`
   } `json:"extensions"`
   Variables struct {
      App string `json:"app"`
      Name string `json:"name"`
      Platform string `json:"platform"`
      Type string `json:"type"`
      UserID string `json:"userId"` // can be empty
   } `json:"variables"`
}

type vodRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID int `json:"accountId"`
   } `json:"mpx"`
}
