package cbc

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "strings"
)

const forwardedFor = "99.224.0.0"

var LogLevel format.LogLevel

// gem.cbc.ca/media/downton-abbey/s01e05
func GetID(addr string) string {
   _, after, _ := strings.Cut(addr, "/media/")
   return after
}

type Asset struct {
   AirDate int64
   Duration int64
   ID string
   PlaySession struct {
      MediaID string
      URL string
   }
   Series string
   Title string
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

type Media struct {
   Message string
   URL string
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
