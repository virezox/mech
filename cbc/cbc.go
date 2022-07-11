package cbc

import (
   "encoding/json"
   "errors"
   "github.com/89z/rosso/http"
   "strings"
   "time"
)

const forwarded_for = "99.224.0.0"

var Client = http.Default_Client

// gem.cbc.ca/media/downton-abbey/s01e05
func Get_ID(input string) string {
   _, after, found := strings.Cut(input, "/media/")
   if found {
      return after
   }
   return input
}

type Asset struct {
   AppleContentId string
   Series string
   Title string
   AirDate int64
   Duration int64
   PlaySession struct {
      URL string
   }
}

func New_Asset(id string) (*Asset, error) {
   var b strings.Builder
   b.WriteString("https://services.radio-canada.ca/ott/cbc-api/v2/assets/")
   b.WriteString(id)
   res, err := Client.Get(b.String())
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

func (a Asset) Get_Duration() time.Duration {
   return time.Duration(a.Duration) * time.Second
}

func (a Asset) Get_Time() time.Time {
   return time.UnixMilli(a.AirDate)
}

func (a Asset) String() string {
   var b strings.Builder
   b.WriteString("ID: ")
   b.WriteString(a.AppleContentId)
   b.WriteString("\nSeries: ")
   b.WriteString(a.Series)
   b.WriteString("\nTitle: ")
   b.WriteString(a.Title)
   b.WriteString("\nDate: ")
   b.WriteString(a.Get_Time().String())
   b.WriteString("\nDuration: ")
   b.WriteString(a.Get_Duration().String())
   return b.String()
}

type Media struct {
   Message *string
   URL *string
}

func (p Profile) Media(asset *Asset) (*Media, error) {
   req, err := http.NewRequest("GET", asset.PlaySession.URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Claims-Token": {p.ClaimsToken},
      "X-Forwarded-For": {forwarded_for},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   if med.Message != nil {
      return nil, errors.New(*med.Message)
   }
   return med, nil
}
