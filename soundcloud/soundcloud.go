package soundcloud

import (
   "encoding/json"
   "fmt"
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
   q := req.URL.Query()
   q.Set("format", "json")
   q.Set("url", addr)
   req.URL.RawQuery = q.Encode()
   fmt.Println("GET", req.URL)
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
   URL string
}

type Track struct {
   ID json.Number
   Title string
   Display_Date string
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         URL string
      }
   }
}

func Resolve(addr string) (*Track, error) {
   req, err := http.NewRequest("GET", Origin + "/resolve", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", clientID)
   q.Set("url", addr)
   req.URL.RawQuery = q.Encode()
   fmt.Println("GET", req.URL)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, fmt.Errorf("status %q", res.Status)
   }
   trk := new(Track)
   if err := json.NewDecoder(res.Body).Decode(trk); err != nil {
      return nil, err
   }
   return trk, nil
}

func Tracks(id string) ([]Track, error) {
   req, err := http.NewRequest("GET", Origin + "/tracks", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", clientID)
   q.Set("ids", id)
   req.URL.RawQuery = q.Encode()
   fmt.Println("GET", req.URL)
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

func (t Track) GetMedia() (*Media, error) {
   var addr string
   for _, code := range t.Media.Transcodings {
      if code.Format.Protocol == "progressive" {
         addr = code.URL
      }
   }
   req, err := http.NewRequest("GET", addr, nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", clientID)
   req.URL.RawQuery = q.Encode()
   fmt.Println("GET", req.URL)
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
