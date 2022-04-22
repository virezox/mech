package cbc

import (
   "bytes"
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strconv"
   "strings"
)

const apiKey = "3f4beddd-2061-49b0-ae80-6f1f2ed65b37"

var LogLevel format.LogLevel

type Asset struct {
   ID string
   PlaySession struct {
      MediaID string
      URL string
   }
}

func NewAsset(addr string) (*Asset, error) {
   _, after, found := strings.Cut(addr, "/media/")
   if !found {
      return nil, notFound{"/media/"}
   }
   req, err := http.NewRequest(
      "GET",
      "https://services.radio-canada.ca/ott/cbc-api/v2/assets/" + after,
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

type Profile struct {
   Tier string
   ClaimsToken string
}

func (p Profile) Media(asset *Asset) (*Media, error) {
   req, err := http.NewRequest("GET", asset.PlaySession.URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Claims-Token": {p.ClaimsToken},
      "X-Forwarded-For": {"99.246.97.250"},
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

type notFound struct {
   value string
}

func (n notFound) Error() string {
   return strconv.Quote(n.value) + " is not found"
}
