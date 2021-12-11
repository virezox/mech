package soundcloud

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "net/url"
)

const (
   Origin = "https://api-v2.soundcloud.com"
   Placeholder = "https://soundcloud.com/images/fb_placeholder.png"
   clientID = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"
)

type Alternate struct {
   Thumbnail_URL string
   Author_URL string
}

func Oembed(addr string) (*Alternate, error) {
   req, err := http.NewRequest("GET", "https://soundcloud.com/oembed", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "format=json&url=" + url.QueryEscape(addr)
   mech.Dump(req)
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
   req, err := http.NewRequest("GET", Origin + "/resolve", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + clientID + "&url=" + url.QueryEscape(addr)
   mech.Dump(req)
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
   req, err := http.NewRequest("GET", Origin + "/tracks", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + clientID + "&ids=" + url.QueryEscape(ids)
   mech.Dump(req)
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
      return nil, mech.NotFound{"progressive"}
   }
   req, err := http.NewRequest("GET", addr + "?client_id=" + clientID, nil)
   if err != nil {
      return nil, err
   }
   mech.Dump(req)
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
