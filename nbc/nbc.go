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

type Bonanza_Page struct {
   Metadata struct {
      MPX_Account_ID string `json:"mpxAccountId"`
      MPX_GUID string `json:"mpxGuid"`
      Series_Short_Title string `json:"seriesShortTitle"`
      Secondary_Title string `json:"secondaryTitle"`
   }
}

func (b Bonanza_Page) Name() string {
   var s strings.Builder
   s.WriteString(b.Metadata.Series_Short_Title)
   s.WriteByte('-')
   s.WriteString(b.Metadata.Secondary_Title)
   return s.String()
}

var (
   Client = http.Default_Client
   secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")
)

func authorization() string {
   now := strconv.FormatInt(time.Now().UnixMilli(), 10)
   str := new(strings.Builder)
   str.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   str.WriteString(",time=")
   str.WriteString(now)
   str.WriteString(",hash=")
   mac := hmac.New(sha256.New, secret_key)
   io.WriteString(mac, now)
   hex.NewEncoder(str).Write(mac.Sum(nil))
   return str.String()
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"`
      Name string `json:"name"` // String cannot represent a non string value
      Platform string `json:"platform"`
      Type string `json:"type"`
      User_ID string `json:"userId"` // can be empty
   } `json:"variables"`
}

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            mpxAccountId
            mpxGuid
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

func New_Bonanza_Page(guid int64) (*Bonanza_Page, error) {
   var p page_request
   p.Variables.App = "nbc"
   p.Variables.Name = strconv.FormatInt(guid, 10)
   p.Variables.Platform = "android"
   p.Variables.Type = "VIDEO"
   p.Query = strings.ReplaceAll(query, "\n", "")
   body, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://friendship.nbc.co/v2/graphql", bytes.NewReader(body),
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

type Video struct {
   ManifestPath string // this is only valid for one minute
}

type video_request struct {
   Device string `json:"device"`
   Device_ID string `json:"deviceId"`
   External_Advertiser_ID string `json:"externalAdvertiserId"`
   MPX struct {
      Account_ID string `json:"accountId"`
   } `json:"mpx"`
}

func (b Bonanza_Page) Video() (*Video, error) {
   var v video_request
   v.Device = "android"
   v.Device_ID = "android"
   v.External_Advertiser_ID = "NBC"
   v.MPX.Account_ID = b.Metadata.MPX_Account_ID
   body, err := json.MarshalIndent(v, "", " ")
   if err != nil {
      return nil, err
   }
   var str strings.Builder
   str.WriteString("http://access-cloudpath.media.nbcuni.com")
   str.WriteString("/access/vod/nbcuniversal/")
   str.WriteString(b.Metadata.MPX_GUID)
   req, err := http.NewRequest("POST", str.String(), bytes.NewReader(body))
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
