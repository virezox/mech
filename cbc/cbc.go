package cbc

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

const forwardedFor = "99.224.0.0"

type Asset struct {
   ID string
   PlaySession struct {
      MediaID string
      URL string
   }
}

type Media struct {
   Message string
   URL string
}

type OverTheTop struct {
   AccessToken string
}

func (o OverTheTop) Profile() (*Profile, error) {
   req, err := http.NewRequest(
      "GET", "https://services.radio-canada.ca/ott/cbc-api/v2/profile", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("OTT-Access-Token", o.AccessToken)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   pro := new(Profile)
   if err := json.NewDecoder(res.Body).Decode(pro); err != nil {
      return nil, err
   }
   return pro, nil
}

func (p Profile) Media(asset *Asset) (*Media, error) {
   req, err := http.NewRequest("GET", asset.PlaySession.URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Claims-Token": {p.ClaimsToken},
      "X-Forwarded-For": {forwardedFor},
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type WebToken struct {
   Signature string
}

func (w WebToken) OverTheTop() (*OverTheTop, error) {
   buf := new(bytes.Buffer)
   err := json.NewEncoder(buf).Encode(map[string]string{
      "jwt": w.Signature,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://services.radio-canada.ca/ott/cbc-api/v2/token", buf,
   )
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   top := new(OverTheTop)
   if err := json.NewDecoder(res.Body).Decode(top); err != nil {
      return nil, err
   }
   return top, nil
}

// gem.cbc.ca/media/downton-abbey/s01e05
func GetID(addr string) string {
   _, after, _ := strings.Cut(addr, "/media/")
   return after
}

func NewAsset(id string) (*Asset, error) {
   req, err := http.NewRequest(
      "GET",
      "https://services.radio-canada.ca/ott/cbc-api/v2/assets/" + id,
      nil,
   )
   if err != nil {
      return nil, err
   }
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   asset := new(Asset)
   if err := json.NewDecoder(res.Body).Decode(asset); err != nil {
      return nil, err
   }
   return asset, nil
}

type Profile struct {
   Tier string
   ClaimsToken string
}

func (p Profile) Create(elem ...string) error {
   return format.Create(p, elem...)
}

func OpenProfile(elem ...string) (*Profile, error) {
   return format.Open[Profile](elem...)
}
