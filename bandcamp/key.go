package bandcamp

import (
   "encoding/json"
   "github.com/89z/mech"
   "net/http"
   "strconv"
   "time"
)

const (
   ApiAlbum = "http://bandcamp.com/api/album/2/info"
   ApiBand = "http://bandcamp.com/api/band/1/info"
   ApiTrack = "http://bandcamp.com/api/track/3/info"
   ApiUrl = "http://bandcamp.com/api/url/2/info"
)

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

type Album struct {
   Large_Art_URL string // 350 x 350
   Release_Date int64
}

// ID to Album. Request uses key.
func NewAlbum(id int) (*Album, error) {
   req, err := http.NewRequest("GET", ApiAlbum, nil)
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
   alb := new(Album)
   if err := json.NewDecoder(res.Body).Decode(alb); err != nil {
      return nil, err
   }
   return alb, nil
}

func (a Album) Unix() time.Time {
   return time.Unix(a.Release_Date, 0)
}

type Band struct {
   Name string
   URL string
}

// ID to Band. Request uses key.
func NewBand(id int) (*Band, error) {
   req, err := http.NewRequest("GET", ApiBand, nil)
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
   ban := new(Band)
   if err := json.NewDecoder(res.Body).Decode(ban); err != nil {
      return nil, err
   }
   return ban, nil
}

type Info struct {
   Album_ID int
   Band_ID int
   Track_ID int
}

// URL to ID. Request uses key.
func NewInfo(addr string) (*Info, error) {
   req, err := http.NewRequest("GET", ApiUrl, nil)
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
   inf := new(Info)
   if err := json.NewDecoder(res.Body).Decode(inf); err != nil {
      return nil, err
   }
   return inf, nil
}

type Track struct {
   Album_ID int
   Streaming_URL string
   Title string
}

// ID to Track. Request uses key.
func NewTrack(id int) (*Track, error) {
   req, err := http.NewRequest("GET", ApiTrack, nil)
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
   tra := new(Track)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}
