package bleep

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "github.com/89z/parse/html"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

const Origin = "https://bleep.com"

func cut(s string, sep byte) (string, string, bool) {
   i := strings.IndexByte(s, sep)
   if i == -1 {
      return s, "", false
   }
   return s[:i], s[i+1:], true
}

func lastCut(s string, sep byte) (string, string, bool) {
   i := strings.LastIndexByte(s, sep)
   if i == -1 {
      return s, "", false
   }
   return s[:i], s[i+1:], true
}

type Date string

// can be either one of these:
//  2001-05-01 00:00:00.0
//  Tue May 01 00:00:00 UTC 2001
func (d Date) Parse() (time.Time, error) {
   value := string(d)
   date, err := time.Parse(time.UnixDate, value)
   if err == nil {
      return date, nil
   }
   return time.Parse("2006-01-02", value)
}

type Meta struct {
   Image string `json:"og:image"`
   Release_Date Date `json:"music:release_date"`
}

// https://bleep.com/release/8728-four-tet-pause
func NewMeta(addr string) (*Meta, error) {
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   met := new(Meta)
   if err := html.NewMap(res.Body).Struct(met); err != nil {
      return nil, err
   }
   return met, nil
}

type Track struct {
   Disc int
   Number int
   ReleaseID int
   Title string
}

// 8728-1-1
func Parse(track string) (*Track, error) {
   var t Track
   _, err := fmt.Sscanf(track, "%v-%v-%v", &t.ReleaseID, &t.Disc, &t.Number)
   if err != nil {
      return nil, err
   }
   return &t, nil
}

// https://bleep.com/release/8728-four-tet-pause
func Release(addr string) ([]Track, error) {
   _, id, ok := lastCut(addr, '/')
   if ! ok {
      return nil, mech.Invalid{addr}
   }
   id, _, ok = cut(id, '-')
   if ! ok {
      return nil, mech.Invalid{addr}
   }
   val := url.Values{
      "id": {id}, "type": {"ReleaseProduct"},
   }
   req, err := http.NewRequest(
      "POST", Origin + "/player/addToPlaylist", strings.NewReader(val.Encode()),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var rel []Track
   if err := json.NewDecoder(res.Body).Decode(&rel); err != nil {
      return nil, err
   }
   return rel, nil
}

func (t Track) Resolve() (string, error) {
   req, err := http.NewRequest(
      "GET", Origin + "/player/resolve/" + t.String(), nil,
   )
   if err != nil {
      return "", err
   }
   res, err := mech.RoundTrip(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   dst, err := io.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   return string(dst), nil
}

func (t Track) String() string {
   return fmt.Sprint(t.ReleaseID, "-", t.Disc, "-", t.Number)
}
