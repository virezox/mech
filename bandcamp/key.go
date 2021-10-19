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
   apiUrl = "http://bandcamp.com/api/url/2/info"
)

// thrjozkaskhjastaurrtygitylpt
// throtaudvinroftignmarkreina
// ullrettkalladrhampa
const key = "veidihundr"

type AlbumInfo struct {
   Large_Art_URL string // 350 x 350
   Release_Date int64
}

// ID to AlbumInfo. Request uses key.
func NewAlbumInfo(id int) (*AlbumInfo, error) {
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
   alb := new(AlbumInfo)
   if err := json.NewDecoder(res.Body).Decode(alb); err != nil {
      return nil, err
   }
   return alb, nil
}

func (a AlbumInfo) Unix() time.Time {
   return time.Unix(a.Release_Date, 0)
}

type BandInfo struct {
   Name string
   URL string
}

// ID to BandInfo. Request uses key.
func NewBandInfo(id int) (*BandInfo, error) {
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
   ban := new(BandInfo)
   if err := json.NewDecoder(res.Body).Decode(ban); err != nil {
      return nil, err
   }
   return ban, nil
}

type TrackInfo struct {
   Album_ID int
   Streaming_URL string
   Title string
}

// ID to TrackInfo. Request uses key.
func NewTrackInfo(id int) (*TrackInfo, error) {
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
   tra := new(TrackInfo)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}

type UrlInfo struct {
   Album_ID int
   Band_ID int
   Track_ID int
}

// URL to UrlInfo. Request uses key.
func NewUrlInfo(addr string) (*UrlInfo, error) {
   req, err := http.NewRequest("GET", apiUrl, nil)
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
   inf := new(UrlInfo)
   if err := json.NewDecoder(res.Body).Decode(inf); err != nil {
      return nil, err
   }
   return inf, nil
}

func (i UrlInfo) Tralbum() (*Tralbum, error) {
   if i.Track_ID != 0 {
      return NewTralbum('t', i.Track_ID)
   }
   return NewTralbum('a', i.Album_ID)
}
