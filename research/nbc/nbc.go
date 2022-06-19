package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "errors"
   "github.com/89z/format"
   "github.com/89z/mech"
   "io"
   "net/http"
   "strconv"
   "strings"
   "time"
)

const persistedQuery = "872a3dffc3ae6cdb3dc69fe3d9a949b539de7b579e95b2942e68d827b1a6ec62"

var (
   LogLevel format.LogLevel
   secretKey = []byte("2b84a073ede61c766e4c0b3f1e656f7f")
)

func writeHash(dst io.Writer, text string, key []byte) {
   mac := hmac.New(sha256.New, key)
   io.WriteString(mac, text)
   hex.NewEncoder(dst).Write(mac.Sum(nil))
}

type BonanzaPage struct {
   Analytics struct {
      ConvivaAssetName string
   }
   Metadata struct {
      MpxAccountID string
   }
}

func NewBonanzaPage(guid int64) (*BonanzaPage, error) {
   var p pageRequest
   p.Extensions.PersistedQuery.Sha256Hash = persistedQuery
   p.Variables.App = "nbc"
   p.Variables.Name = strconv.FormatInt(guid, 10)
   p.Variables.OneApp = true
   p.Variables.Platform = "android"
   p.Variables.Type = "VIDEO"
   buf, err := mech.Encode(p)
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
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var page struct {
      Data struct {
         BonanzaPage BonanzaPage
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
      return nil, err
   }
   return &page.Data.BonanzaPage, nil
}

func (b BonanzaPage) Base() string {
   return mech.Clean(b.Analytics.ConvivaAssetName)
}

func (b BonanzaPage) Video(guid int64) (*Video, error) {
   var in videoRequest
   in.Device = "android"
   in.DeviceID = "android"
   in.ExternalAdvertiserID = "NBC"
   in.Mpx.AccountID = b.Metadata.MpxAccountID
   body, err := mech.Encode(in)
   if err != nil {
      return nil, err
   }
   addr := []byte("http://access-cloudpath.media.nbcuni.com")
   addr = append(addr, "/access/vod/nbcuniversal/"...)
   addr = strconv.AppendInt(addr, guid, 10)
   req, err := http.NewRequest("POST", string(addr), body)
   if err != nil {
      return nil, err
   }
   unix := strconv.FormatInt(time.Now().UnixMilli(), 10)
   auth := new(strings.Builder)
   auth.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   auth.WriteString(",time=")
   auth.WriteString(unix)
   auth.WriteString(",hash=")
   writeHash(auth, unix, secretKey)
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
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   out := new(Video)
   if err := json.NewDecoder(res.Body).Decode(out); err != nil {
      return nil, err
   }
   return out, nil
}

type Video struct {
   ManifestPath string // this is only valid for one minute
}

type pageRequest struct {
   Extensions struct {
      PersistedQuery struct {
         Sha256Hash string `json:"sha256Hash"`
      } `json:"persistedQuery"`
   } `json:"extensions"`
   Variables struct {
      App string `json:"app"`
      Name string `json:"name"` // String cannot represent a non string value
      OneApp bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"`
      UserID string `json:"userId"` // can be empty
   } `json:"variables"`
}

type videoRequest struct {
   Device string `json:"device"`
   DeviceID string `json:"deviceId"`
   ExternalAdvertiserID string `json:"externalAdvertiserId"`
   Mpx struct {
      AccountID string `json:"accountId"`
   } `json:"mpx"`
}
