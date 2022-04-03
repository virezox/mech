package soundcloud

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Track struct {
   Artwork_URL string
   Display_Date string
   ID int64
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         URL string
      }
   }
   Title string
   User struct {
      Avatar_URL string
      Username string
   }
}

func (t Track) Base() string {
   return t.User.Username + "-" + t.Title
}

func NewTrack(id int64) (*Track, error) {
   buf := []byte("https://api-v2.soundcloud.com/tracks/")
   buf = strconv.AppendInt(buf, id, 10)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + clientID
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

func Resolve(addr string) ([]Track, error) {
   req, err := http.NewRequest(
      "GET", "https://api-v2.soundcloud.com/resolve", nil,
   )
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
   var solve struct {
      Kind string
      Track
   }
   if err := json.NewDecoder(res.Body).Decode(&solve); err != nil {
      return nil, err
   }
   if solve.Kind == "track" {
      return []Track{solve.Track}, nil
   }
   return UserTracks(solve.ID)
}

// We can also paginate, but for now this is good enough.
func UserTracks(id int64) ([]Track, error) {
   buf := []byte("https://api-v2.soundcloud.com/users/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, "/tracks"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {clientID},
      "limit": {"999"},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var user struct {
      Collection []Track
   }
   if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
      return nil, err
   }
   return user.Collection, nil
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (t Track) Artwork() string {
   if t.Artwork_URL == "" {
      t.Artwork_URL = t.User.Avatar_URL
   }
   return strings.Replace(t.Artwork_URL, "large", "t500x", 1)
}

// Also available is "hls", but all transcodings are quality "sq".
// Same for "api-mobile.soundcloud.com".
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
