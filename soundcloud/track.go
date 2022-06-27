package soundcloud
// github.com/89z

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

// Also available is "hls", but all transcodings are quality "sq".
// Same for "api-mobile.soundcloud.com".
func (t Track) Progressive() (*Media, error) {
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
   req.URL.RawQuery = "client_id=" + client_ID
   res, err := Client.Do(req)
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

type Track struct {
   ID int64
   Display_Date string // 2021-04-12T07:00:01Z
   User struct {
      Username string
      Avatar_URL string
   }
   Title string
   Artwork_URL string
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         URL string
      }
   }
}

func New_Track(id int64) (*Track, error) {
   buf := []byte("https://api-v2.soundcloud.com/tracks/")
   buf = strconv.AppendInt(buf, id, 10)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + client_ID
   res, err := Client.Do(req)
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
      "client_id": {client_ID},
      "url": {addr},
   }.Encode()
   res, err := Client.Do(req)
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
   return User_Tracks(solve.ID)
}

// We can also paginate, but for now this is good enough.
func User_Tracks(id int64) ([]Track, error) {
   buf := []byte("https://api-v2.soundcloud.com/users/")
   buf = strconv.AppendInt(buf, id, 10)
   buf = append(buf, "/tracks"...)
   req, err := http.NewRequest("GET", string(buf), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {client_ID},
      "limit": {"999"},
   }.Encode()
   res, err := Client.Do(req)
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

func (t Track) Base() string {
   return t.User.Username + "-" + t.Title
}

func (t Track) String() string {
   var buf []byte
   buf = append(buf, "ID: "...)
   buf = strconv.AppendInt(buf, t.ID, 10)
   buf = append(buf, "\nDate: "...)
   buf = append(buf, t.Display_Date...)
   buf = append(buf, "\nUser: "...)
   buf = append(buf, t.User.Username...)
   buf = append(buf, "\nAvatar: "...)
   buf = append(buf, t.User.Avatar_URL...)
   buf = append(buf, "\nTitle: "...)
   buf = append(buf, t.Title...)
   if t.Artwork_URL != "" {
      buf = append(buf, "\nArtwork: "...)
      buf = append(buf, t.Artwork_URL...)
   }
   for _, coding := range t.Media.Transcodings {
      buf = append(buf, "\nFormat: "...)
      buf = append(buf, coding.Format.Protocol...)
      buf = append(buf, "\nURL: "...)
      buf = append(buf, coding.URL...)
   }
   return string(buf)
}

func (t Track) Time() (time.Time, error) {
   return time.Parse(time.RFC3339, t.Display_Date)
}
