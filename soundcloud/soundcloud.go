package soundcloud

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech"
   "net/http"
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
   val := req.URL.Query()
   val.Set("format", "json")
   val.Set("url", addr)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
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
   val := req.URL.Query()
   val.Set("client_id", clientID)
   val.Set("url", addr)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   trk := new(Track)
   if err := json.NewDecoder(res.Body).Decode(trk); err != nil {
      return nil, err
   }
   return trk, nil
}

func Tracks(ids string) ([]Track, error) {
   req, err := http.NewRequest("GET", Origin + "/tracks", nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("client_id", clientID)
   val.Set("ids", ids)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
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

func (t Track) progressive() string {
   for _, code := range t.Media.Transcodings {
      if code.Format.Protocol == "progressive" {
         return code.URL
      }
   }
   return ""
}

func (t Track) GetMedia() (*Media, error) {
   addr := t.progressive()
   if addr == "" {
      return nil, fmt.Errorf("%+v", t)
   }
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("client_id", clientID)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
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
