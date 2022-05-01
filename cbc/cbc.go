package cbc

import (
   "encoding/json"
   "fmt"
   "github.com/89z/format"
   "net/http"
   "strings"
   "time"
)

const forwardedFor = "99.224.0.0"

var LogLevel format.LogLevel

// gem.cbc.ca/media/downton-abbey/s01e05
func GetID(addr string) string {
   _, after, _ := strings.Cut(addr, "/media/")
   return after
}

type Asset struct {
   AppleContentID string
   Series string
   Title string
   AirDate int64
   Duration int64
   PlaySession struct {
      URL string
   }
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

func (a Asset) Format(f fmt.State, verb rune) {
   fmt.Fprintln(f, "ID:", a.AppleContentID)
   fmt.Fprintln(f, "Series:", a.Series)
   fmt.Fprintln(f, "Title:", a.Title)
   fmt.Fprintln(f, "Date:", a.GetTime())
   fmt.Fprint(f, "Duration: ", a.GetDuration())
   if verb == 'a' {
      fmt.Fprint(f, "\nURL: ", a.PlaySession.URL)
   }
}

func (a Asset) GetDuration() time.Duration {
   return time.Duration(a.Duration) * time.Second
}

func (a Asset) GetTime() time.Time {
   return time.UnixMilli(a.AirDate)
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
