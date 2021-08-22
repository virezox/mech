package soundcloud

import (
   "encoding/json"
   "net/http"
   "net/http/httputil"
   "os"
)

const (
   Origin = "https://api-v2.soundcloud.com"
   clientID = "fSSdm5yTnDka1g0Fz1CO5Yx6z0NbeHAj"
)

// MediaURLResponse is the JSON response of retrieving media information of a
// track
type Media struct {
   URL string
}

type Track struct {
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

func Tracks(id string) ([]Track, error) {
   req, err := http.NewRequest("GET", Origin + "/tracks", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", clientID)
   q.Set("ids", id)
   req.URL.RawQuery = q.Encode()
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
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

func Resolve(addr string) (*Track, error) {
   req, err := http.NewRequest("GET", Origin + "/resolve", nil)
   if err != nil {
      return nil, err
   }
   q := req.URL.Query()
   q.Set("client_id", clientID)
   q.Set("url", addr)
   req.URL.RawQuery = q.Encode()
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   t := new(Track)
   if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
      return nil, err
   }
   return t, nil
}

// The media URL is the actual link to the audio file for the track. "addr" is
// Track.Media.Transcodings[0].URL
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
   d, err := httputil.DumpRequest(req, false)
   if err != nil {
      return nil, err
   }
   os.Stdout.Write(d)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   m := new(Media)
   if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
      return nil, err
   }
   return m, nil
}
