package soundcloud

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

var ImageSizes = []string{
   "t120x120",
   "t1240x260",
   "t200x200",
   "t20x20",
   "t240x240",
   "t2480x520",
   "t250x250",
   "t300x300",
   "t40x40",
   "t47x47",
   "t500x",
   "t500x500",
   "t50x50",
   "t60x60",
   "t67x67",
   "t80x80",
   "tx250",
}

const (
   clientID = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"
   origin = "https://api-v2.soundcloud.com"
)

var LogLevel format.LogLevel

type Media struct {
   // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJ...
   URL string
}

type Track struct {
   Display_Date string // 2021-04-12T07:00:01Z
   ID int
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         // api-v2.soundcloud.com/media/soundcloud:tracks:103650107/
         // aca81dd5-2feb-4fc4-a102-036fb35fe44a/stream/progressive
         URL string
      }
   }
   Title string
   Artwork_URL string
   User struct {
      Avatar_URL string
      Username string
   }
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (t Track) Artwork() string {
   if t.Artwork_URL == "" {
      t.Artwork_URL = t.User.Avatar_URL
   }
   return strings.Replace(t.Artwork_URL, "large", "t500x500", 1)
}

func Resolve(addr string) (*Track, error) {
   req, err := http.NewRequest("GET", origin + "/resolve", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {clientID},
      "url": {addr},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(Track)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}

func Tracks(ids []int64) ([]Track, error) {
   req, err := http.NewRequest("GET", origin + "/tracks", nil)
   if err != nil {
      return nil, err
   }
   var buf []byte
   for key, val := range ids {
      if key >= 1 {
         buf = append(buf, ',')
      }
      buf = strconv.AppendInt(buf, val, 10)
   }
   req.URL.RawQuery = url.Values{
      "client_id": {clientID},
      "ids": {string(buf)},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var tracks []Track
   if err := json.NewDecoder(res.Body).Decode(&tracks); err != nil {
      return nil, err
   }
   return tracks, nil
}

func (t Track) Progressive() (*Media, error) {
   var addr string
   for _, code := range t.Media.Transcodings {
      if code.Format.Protocol == "progressive" {
         addr = code.URL
      }
   }
   req, err := http.NewRequest("GET", addr + "?client_id=" + clientID, nil)
   if err != nil {
      return nil, err
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
