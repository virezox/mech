package bandcamp

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "time"
)

const (
   apiAlbum = "http://bandcamp.com/api/album/2/info"
   apiBand = "http://bandcamp.com/api/band/1/info"
   apiTrack = "http://bandcamp.com/api/track/3/info"
   apiURL = "http://bandcamp.com/api/url/2/info"
)

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

type InfoAlbum struct {
   Large_Art_URL string // 350 x 350
   Release_Date int64
}

// ID to InfoAlbum. Request uses key.
func NewInfoAlbum(id int) (*InfoAlbum, error) {
   req, err := http.NewRequest("GET", apiAlbum, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("album_id", strconv.Itoa(id))
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   alb := new(InfoAlbum)
   if err := json.NewDecoder(res.Body).Decode(alb); err != nil {
      return nil, err
   }
   return alb, nil
}

func (a InfoAlbum) Unix() time.Time {
   return time.Unix(a.Release_Date, 0)
}

type InfoBand struct {
   Name string
   URL string
}

// ID to InfoBand. Request uses key.
func NewInfoBand(id int) (*InfoBand, error) {
   req, err := http.NewRequest("GET", apiBand, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("band_id", strconv.Itoa(id))
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   ban := new(InfoBand)
   if err := json.NewDecoder(res.Body).Decode(ban); err != nil {
      return nil, err
   }
   return ban, nil
}

type InfoTrack struct {
   Album_ID int
   Streaming_URL string
   Title string
}

// ID to InfoTrack. Request uses key.
func NewInfoTrack(id int) (*InfoTrack, error) {
   req, err := http.NewRequest("GET", apiTrack, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("track_id", strconv.Itoa(id))
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(InfoTrack)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}

type InfoURL struct {
   Album_ID int
   Band_ID int
   Track_ID int
}

// URL to InfoURL. Request uses key.
func NewInfoURL(addr string) (*InfoURL, error) {
   req, err := http.NewRequest("GET", apiURL, nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("key", key)
   val.Set("url", addr)
   req.URL.RawQuery = val.Encode()
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   inf := new(InfoURL)
   if err := json.NewDecoder(res.Body).Decode(inf); err != nil {
      return nil, err
   }
   return inf, nil
}
