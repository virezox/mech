package soundcloud

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

const clientID = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"

var LogLevel format.LogLevel

type Media struct {
   // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJ...
   URL string
}

type Track struct {
   Artwork_URL_Template string
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
   User struct {
      Username string
   }
}

func Resolve(addr string) (*Track, error) {
   req, err := http.NewRequest(
      "GET", "https://api-v2.soundcloud.com/resolve", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {clientID}, "url": {addr},
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
   req, err := http.NewRequest(
      "GET", "https://api-v2.soundcloud.com/tracks", nil,
   )
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

// i1.sndcdn.com/avatars-000274827119-0dxutu-{size}.jpg
// i1.sndcdn.com/avatars-000274827119-0dxutu-t500x500.jpg
func (t Track) Artwork() string {
   return strings.Replace(t.Artwork_URL_Template, "{size}", "t500x500", 1)
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
