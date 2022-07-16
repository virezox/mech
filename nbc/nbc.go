package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "github.com/89z/rosso/http"
   "io"
   "strconv"
   "strings"
   "time"
)

func (b Bonanza_Page) Video() (*Video, error) {
   var v video_request
   v.Device = "android"
   v.Device_ID = "android"
   v.External_Advertiser_ID = "NBC"
   v.Mpx.Account_ID = b.Metadata.MpxAccountId
   body, err := json.MarshalIndent(v, "", " ")
   if err != nil {
      return nil, err
   }
   var buf strings.Builder
   buf.WriteString("http://access-cloudpath.media.nbcuni.com")
   buf.WriteString("/access/vod/nbcuniversal/")
   buf.WriteString(b.Name)
   req, err := http.NewRequest("POST", buf.String(), bytes.NewReader(body))
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {authorization()},
      "Content-Type": {"application/json"},
   }
   res, err := Client.Do(req)
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

type Video struct {
   ManifestPath string // this is only valid for one minute
}

type page_request struct {
   Extensions struct {
      Persisted_Query struct {
         SHA_256_Hash string `json:"sha256Hash"`
      } `json:"persistedQuery"`
   } `json:"extensions"`
   Variables struct {
      App string `json:"app"`
      Name string `json:"name"` // String cannot represent a non string value
      One_App bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"`
      User_ID string `json:"userId"` // can be empty
   } `json:"variables"`
}

type video_request struct {
   Device string `json:"device"`
   Device_ID string `json:"deviceId"`
   External_Advertiser_ID string `json:"externalAdvertiserId"`
   Mpx struct {
      Account_ID string `json:"accountId"`
   } `json:"mpx"`
}
const persisted_query = "872a3dffc3ae6cdb3dc69fe3d9a949b539de7b579e95b2942e68d827b1a6ec62"

var (
   Client = http.Default_Client
   secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")
)

func authorization() string {
   now := strconv.FormatInt(time.Now().UnixMilli(), 10)
   b := new(strings.Builder)
   b.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   b.WriteString(",time=")
   b.WriteString(now)
   b.WriteString(",hash=")
   mac := hmac.New(sha256.New, secret_key)
   io.WriteString(mac, now)
   hex.NewEncoder(b).Write(mac.Sum(nil))
   return b.String()
}

func New_Bonanza_Page(guid int64) (*Bonanza_Page, error) {
   var p page_request
   p.Extensions.Persisted_Query.SHA_256_Hash = persisted_query
   p.Variables.App = "nbc"
   p.Variables.Name = strconv.FormatInt(guid, 10)
   p.Variables.One_App = true
   p.Variables.Platform = "android"
   p.Variables.Type = "VIDEO"
   buf, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://friendship.nbc.co/v2/graphql", bytes.NewReader(buf),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var page struct {
      Data struct {
         BonanzaPage Bonanza_Page
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
      return nil, err
   }
   return &page.Data.BonanzaPage, nil
}

type Bonanza_Page struct {
   Analytics struct {
      ConvivaAssetName string
   }
   Metadata struct {
      MpxAccountId string
   }
   Name string
}

