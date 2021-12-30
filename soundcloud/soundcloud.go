package soundcloud

import (
   "encoding/json"
   "github.com/89z/format"
   "net/http"
   "net/url"
   "strconv"
)

const (
   Placeholder = "https://soundcloud.com/images/fb_placeholder.png"
   clientID = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"
   origin = "https://api-v2.soundcloud.com"
)

var LogLevel format.LogLevel

type Alternate struct {
   Thumbnail_URL string
   Author_URL string
}

func Oembed(addr string) (*Alternate, error) {
   req, err := http.NewRequest("GET", "https://soundcloud.com/oembed", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "format": {"json"}, "url": {addr},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   emb := new(Alternate)
   if err := json.NewDecoder(res.Body).Decode(emb); err != nil {
      return nil, err
   }
   return emb, nil
}

type Media struct {
   // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJl...
   URL string
}

type Track struct {
   ID int
   Title string
   // 2021-04-12T07:00:01Z
   Display_Date string
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
   User struct {
      Username string
   }
}

func Resolve(addr string) (*Track, error) {
   req, err := http.NewRequest("GET", origin + "/resolve", nil)
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

func Tracks(ids string) ([]Track, error) {
   req, err := http.NewRequest("GET", origin + "/tracks", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {clientID}, "ids": {ids},
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
   if addr == "" {
      return nil, notFound{"progressive"}
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

type notFound struct {
   input string
}

func (n notFound) Error() string {
   return strconv.Quote(n.input) + " not found"
}
